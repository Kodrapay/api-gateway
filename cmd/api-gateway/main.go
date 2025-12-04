package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kodra-pay/api-gateway/internal/config"
	"github.com/kodra-pay/api-gateway/internal/middleware"
	"github.com/kodra-pay/api-gateway/internal/routes"
)

func main() {
	cfg := config.Load("api-gateway", "7000")

	app := fiber.New()
	app.Use(middleware.RequestID())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "*",
		AllowCredentials: true,
	}))
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodOptions {
			c.Set("Access-Control-Allow-Origin", "*")
			c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Set("Access-Control-Allow-Headers", "*")
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})

	routes.Register(app, cfg.ServiceName)

	log.Printf("%s listening on :%s", cfg.ServiceName, cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
