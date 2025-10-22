package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		latency := time.Since(start)
		method := c.Method()
		path := c.OriginalURL()
		status := c.Response().StatusCode()

		log.Printf("[%d] %s %s (%s)", status, method, path, latency)

		return err
	}
}
