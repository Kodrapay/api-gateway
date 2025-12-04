package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type GatewayHandler struct{}

func NewGatewayHandler() *GatewayHandler { return &GatewayHandler{} }

// ProxyRequest forwards the request to the appropriate backend service
func (h *GatewayHandler) ProxyRequest(serviceURL string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Build the target URL
		path := c.OriginalURL()
		queryString := string(c.Request().URI().QueryString())
		targetURL := fmt.Sprintf("http://%s%s", serviceURL, path)
		if queryString != "" {
			targetURL += "?" + queryString
		}

		// Get the raw body bytes before any parsing
		bodyBytes := c.Request().Body()
		req, err := http.NewRequest(string(c.Method()), targetURL, bytes.NewReader(bodyBytes))
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "failed to create proxy request"})
		}

		// Copy headers
		c.Request().Header.VisitAll(func(key, value []byte) {
			if string(key) != "Host" {
				req.Header.Set(string(key), string(value))
			}
		})

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "service unavailable"})
		}
		defer resp.Body.Close()

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to read response"})
		}

		// Copy response headers, avoid duplicating CORS headers (handled by gateway middleware)
		for k, v := range resp.Header {
			if k == "Access-Control-Allow-Origin" || k == "Access-Control-Allow-Methods" || k == "Access-Control-Allow-Headers" {
				continue
			}
			if len(v) > 0 {
				c.Response().Header.Set(k, v[0])
			}
		}

		return c.Status(resp.StatusCode).Send(body)
	}
}

func (h *GatewayHandler) Routes(c *fiber.Ctx) error {
	routes := fiber.Map{
		"api_gateway": fiber.Map{
			"port":   7000,
			"status": "implemented",
		},
		"auth_service": fiber.Map{
			"port":      7001,
			"endpoints": []string{"/auth/login", "/auth/register", "/auth/verify", "/auth/refresh"},
			"status":    "pending",
		},
		"merchant_service": fiber.Map{
			"port":      7002,
			"endpoints": []string{"/merchants", "/merchants/:id", "/merchants/profile", "/merchants/settings"},
			"status":    "pending",
		},
		"admin_service": fiber.Map{
			"port":      7003,
			"endpoints": []string{"/admin/users", "/admin/merchants", "/admin/transactions", "/admin/reports"},
			"status":    "pending",
		},
		"transaction_service": fiber.Map{
			"port":      7004,
			"endpoints": []string{"/transactions", "/transactions/:id", "/transactions/search"},
			"status":    "pending",
		},
		"checkout_service": fiber.Map{
			"port":      7005,
			"endpoints": []string{"/checkout", "/checkout/:id", "/checkout/process"},
			"status":    "pending",
		},
		"webhook_service": fiber.Map{
			"port":      7006,
			"endpoints": []string{"/webhooks", "/webhooks/:id", "/webhooks/logs"},
			"status":    "pending",
		},
		"wallet_ledger_service": fiber.Map{
			"port":      7007,
			"endpoints": []string{"/wallets", "/wallets/:id", "/wallets/ledger"},
			"status":    "pending",
		},
		"settlement_service": fiber.Map{
			"port":      7008,
			"endpoints": []string{"/settlements", "/settlements/:id", "/settlements/reconcile"},
			"status":    "pending",
		},
		"payout_service": fiber.Map{
			"port":      7009,
			"endpoints": []string{"/payouts", "/payouts/:id", "/payouts/schedule"},
			"status":    "pending",
		},
		"virtual_account_service": fiber.Map{
			"port":      7010,
			"endpoints": []string{"/accounts", "/accounts/:id", "/accounts/transactions"},
			"status":    "pending",
		},
		"reconciliation_service": fiber.Map{
			"port":      7011,
			"endpoints": []string{"/reconciliation", "/reconciliation/:id", "/reconciliation/process"},
			"status":    "pending",
		},
		"fraud_service": fiber.Map{
			"port":      7012,
			"endpoints": []string{"/fraud/check", "/fraud/rules", "/fraud/alerts"},
			"status":    "pending",
		},
		"dispute_service": fiber.Map{
			"port":      7013,
			"endpoints": []string{"/disputes", "/disputes/:id", "/disputes/resolve"},
			"status":    "pending",
		},
		"notification_service": fiber.Map{
			"port":      7014,
			"endpoints": []string{"/notifications", "/notifications/send", "/notifications/templates"},
			"status":    "pending",
		},
		"compliance_service": fiber.Map{
			"port":      7015,
			"endpoints": []string{"/compliance/kyc", "/compliance/aml", "/compliance/reports"},
			"status":    "pending",
		},
		"encryption_service": fiber.Map{
			"port":      7016,
			"endpoints": []string{"/encrypt", "/decrypt", "/keys"},
			"status":    "pending",
		},
		"fee_service": fiber.Map{
			"port":      7017,
			"endpoints": []string{"/fees", "/fees/calculate", "/fees/rules"},
			"status":    "pending",
		},
		"fx_service": fiber.Map{
			"port":      7018,
			"endpoints": []string{"/rates", "/rates/convert", "/rates/historical"},
			"status":    "pending",
		},
	}
	return c.JSON(routes)
}
