package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kodra-pay/api-gateway/internal/config"
	"github.com/kodra-pay/api-gateway/internal/middleware"
	"github.com/kodra-pay/api-gateway/internal/routes"
)

func main() {
	cfg := config.Load("api-gateway", "8000")

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middleware.RequestID())

	// Enable CORS for frontend dashboards
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173,http://localhost:5174,http://127.0.0.1:5173,http://127.0.0.1:5174",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-Request-ID",
		AllowCredentials: true,
	}))

	routes.Register(app, cfg.ServiceName)

	log.Printf("%s listening on :%s", cfg.ServiceName, cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
