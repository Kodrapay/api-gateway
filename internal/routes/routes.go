package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodra-pay/api-gateway/internal/handlers"
)

func Register(app *fiber.App, service string) {
	health := handlers.NewHealthHandler(service)
	health.Register(app)

	h := handlers.NewGatewayHandler()
	app.Get("/routes", h.Routes)
}
