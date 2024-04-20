package middleware

import (
	"github.com/IshaqNiloy/go-rest-api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func AuthReq() func(*fiber.Ctx) error {
	// Configure basic authentication
	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("USERNAME"): config.Config("PASSWORD"),
		},
	}

	// Create a new middleware instance for basic authentication
	authMiddleware := basicauth.New(cfg)

	// Define and return the middleware function
	return func(c *fiber.Ctx) error {
		// Execute the basic authentication middleware
		err := authMiddleware(c)

		// Handle any authentication errors
		if err != nil {
			// Return a 401 Unauthorized response
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Proceed to the next middleware or route handler
		return c.Next()
	}
}
