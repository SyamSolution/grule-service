package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SyamSolution/grule-service/config"
	"github.com/SyamSolution/grule-service/internal/handler"
	"github.com/SyamSolution/grule-service/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gorules/zen-go"
	"github.com/joho/godotenv"
)

func main() {
	baseDep := config.NewBaseDep()
	loadEnv(baseDep.Logger)

	//prometheus.MustRegister(dbCollector)
	//fiberProm := middleware.NewWithRegistry(prometheus.DefaultRegisterer, "transaction-service", "", "", map[string]string{})

	rulebaseEligible := loadGoruleBase("./config/eligible-graph.json")
	rulebaseDiscount := loadGoruleBase("./config/discount-graph.json")

	//=== usecase lists start ===//
	eligibleUsecase := usecase.NewEligibleUsecase(baseDep.Logger, rulebaseEligible)
	discountUsecase := usecase.NewDiscountUsecase(baseDep.Logger, rulebaseDiscount)
	//=== usecase lists end ===//

	//=== handler lists start ===//
	eligibleHandler := handler.NewEligibleHandler(eligibleUsecase, baseDep.Logger)
	discountHandler := handler.NewDiscountHandler(discountUsecase, baseDep.Logger)
	//=== handler lists end ===//

	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(pprof.New())
	app.Use(logger.New(logger.Config{
		// Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
		Format:       `${time} {"router_activity" : [${status},"${latency}","${method}","${path}"], "query_param":${queryParams}, "body_param":${body}}` + "\n",
		TimeInterval: time.Millisecond,
		TimeFormat:   "02-01-2006 15:04:05",
		TimeZone:     "Indonesia/Jakarta",
	}))

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

func loadGoruleBase(ruleFile string) zen.Decision {

	engine := zen.NewEngine(zen.EngineConfig{})

	graph, err := os.ReadFile(ruleFile)
	if err != nil {
		panic(err)
	}

	decision, err := engine.CreateDecision(graph)
	if err != nil {
		panic(err)
	}
	return decision
}
