package middleware

import (
	"github.com/SyamSolution/transaction-service/helper"
	"github.com/gofiber/fiber/v2"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		email, err := helper.VerifyToken(tokenString, "email")
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		c.Locals("email", email)
		return c.Next()
	}
}

func SecureHeadersMiddleware(c *fiber.Ctx) error {
	c.Set("Content-Security-Policy", "default-src 'self'")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Frame-Options", "DENY")
	c.Set("X-XSS-Protection", "1; mode=block")

	return c.Next()
}