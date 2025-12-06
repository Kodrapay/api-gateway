package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodra-pay/api-gateway/internal/handlers"
)

func Register(app *fiber.App, serviceName string) {
	health := handlers.NewHealthHandler(serviceName)
	health.Register(app)

	gateway := handlers.NewGatewayHandler()

	// Register routes handler
	app.Get("/routes", gateway.Routes)

	// Auth Service (Port 7001)
	app.All("/auth/*", gateway.ProxyRequest("auth-service:7001"))

	// Merchant Service (Port 7002)
	app.All("/merchants/*", gateway.ProxyRequest("merchant-service:7002"))
	app.All("/merchant/me", gateway.ProxyRequest("merchant-service:7002"))
	app.All("/payment-links/:id", gateway.ProxyRequest("merchant-service:7002"))
	app.All("/payment-links", gateway.ProxyRequest("merchant-service:7002"))
	// Backward compatibility for singular path
	app.All("/merchant/*", gateway.ProxyRequest("merchant-service:7002"))

	// Admin Service (Port 7003)
	app.All("/admin/*", gateway.ProxyRequest("admin-service:7003"))

	// Transaction Service (Port 7004)
	app.All("/transactions", gateway.ProxyRequest("transaction-service:7004"))
	app.All("/transactions/*", gateway.ProxyRequest("transaction-service:7004"))

	// Checkout Service (Port 7005)
	app.All("/checkout/*", gateway.ProxyRequest("checkout-service:7005"))

	// Webhook Service (Port 7006)
	app.All("/webhooks/*", gateway.ProxyRequest("webhook-service:7006"))

	// Wallet Ledger Service (Port 7007)
	app.All("/wallets/*", gateway.ProxyRequest("wallet-ledger-service:7007"))

	// Settlement Service (Port 7008)
	app.All("/settlements/*", gateway.ProxyRequest("settlement-service:7008"))

	// Payout Service (Port 7009)
	app.All("/payouts", gateway.ProxyRequest("payout-service:7009"))
	app.All("/payouts/*", gateway.ProxyRequest("payout-service:7009"))

	// Virtual Account Service (Port 7010)
	app.All("/accounts/*", gateway.ProxyRequest("virtual-account-service:7010"))

	// Reconciliation Service (Port 7011)
	app.All("/reconciliation/*", gateway.ProxyRequest("reconciliation-service:7011"))

	// Fraud Service (Port 7012)
	app.All("/fraud/*", gateway.ProxyRequest("fraud-service:7012"))

	// Dispute Service (Port 7013)
	app.All("/disputes/*", gateway.ProxyRequest("dispute-service:7013"))
	app.All("/disputes", gateway.ProxyRequest("dispute-service:7013"))

	// Notification Service (Port 7014)
	app.All("/notifications/*", gateway.ProxyRequest("notification-service:7014"))

	// Compliance Service (Port 7015) - KYC and AML
	app.All("/compliance/*", gateway.ProxyRequest("compliance-service:7015"))
	// Route KYC requests to compliance service
	app.All("/kyc/*", gateway.ProxyRequest("compliance-service:7015"))
	app.All("/kyc", gateway.ProxyRequest("compliance-service:7015"))

	// Encryption Service (Port 7016)
	app.All("/encrypt/*", gateway.ProxyRequest("encryption-service:7016"))
	app.All("/decrypt/*", gateway.ProxyRequest("encryption-service:7016"))
	app.All("/keys/*", gateway.ProxyRequest("encryption-service:7016"))

	// Fee Service (Port 7017)
	app.All("/fees/*", gateway.ProxyRequest("fee-service:7017"))

	// FX Service (Port 7018)
	app.All("/rates/*", gateway.ProxyRequest("fx-service:7018"))
}
