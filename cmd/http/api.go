package main

import (
	"fmt"
	"github.com/SyamSolution/grule-service/config"
	"github.com/SyamSolution/grule-service/internal/handler"
	"github.com/SyamSolution/grule-service/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	baseDep := config.NewBaseDep()
	loadEnv(baseDep.Logger)

	//=== usecase lists start ===//
	eligibleUsecase := usecase.NewEligibleUsecase(baseDep.Logger)
	//=== usecase lists end ===//

	//=== handler lists start ===//
	eligibleHandler := handler.NewEligibleHandler(eligibleUsecase, baseDep.Logger)
	//=== handler lists end ===//

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//=== transaction routes ===//
	app.Post("/check-eligible", eligibleHandler.CheckEligibility)

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
