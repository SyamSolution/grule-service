package main

import (
	"fmt"
	"github.com/SyamSolution/grule-service/config"
	"github.com/SyamSolution/grule-service/internal/handler"
	"github.com/SyamSolution/grule-service/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	baseDep := config.NewBaseDep()
	loadEnv(baseDep.Logger)

	//prometheus.MustRegister(dbCollector)
	//fiberProm := middleware.NewWithRegistry(prometheus.DefaultRegisterer, "transaction-service", "", "", map[string]string{})

	knowledgeEligibleBase := knowledgeBase("eligible", "EligibleRules")
	knowledgeDiscountBase := knowledgeBase("discount", "DiscountRules")

	//=== usecase lists start ===//
	eligibleUsecase := usecase.NewEligibleUsecase(baseDep.Logger, knowledgeEligibleBase)
	discountUsecase := usecase.NewDiscountUsecase(baseDep.Logger, knowledgeDiscountBase)
	//=== usecase lists end ===//

	//=== handler lists start ===//
	eligibleHandler := handler.NewEligibleHandler(eligibleUsecase, baseDep.Logger)
	discountHandler := handler.NewDiscountHandler(discountUsecase, baseDep.Logger)
	//=== handler lists end ===//

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//=== rules routes ===//
	app.Post("/check-eligible", eligibleHandler.CheckEligibility)
	app.Post("/check-discount", discountHandler.CheckDiscount)

	//=== listen port ===//
	if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Fatal(err)
	}
}

func loadEnv(logger config.Logger) {
	_, err := os.Stat(".env")
	if err == nil {
		err = godotenv.Load()
		if err != nil {
			logger.Error("no .env files provided")
		}
	}
}

func knowledgeBase(ruleFile, ruleName string) *ast.KnowledgeBase {
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	fileRes := pkg.NewFileResource("./config/" + ruleFile + ".grl")
	err := ruleBuilder.BuildRuleFromResource(ruleName, "0.0.1", fileRes)
	if err != nil {
		panic(err)
	}

	knowledgeBase, _ := knowledgeLibrary.NewKnowledgeBaseInstance(ruleName, "0.0.1")

	return knowledgeBase
}
