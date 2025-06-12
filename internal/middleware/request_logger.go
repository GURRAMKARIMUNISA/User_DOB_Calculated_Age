package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/GurramKarimunisa/go-user-api/internal/logger" // Adjust import path
	"go.uber.org/zap"
)

// RequestLogger logs incoming requests and their duration
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		// Call StatusCode() to get the integer value
		statusCode := c.Response().StatusCode()
		method := c.Method()
		path := c.Path()
		ip := c.IP()

		if err != nil {
			logger.Log.Error(
				"Request failed",
				zap.Int("status", statusCode),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("ip", ip),
				zap.Duration("duration", duration),
				zap.Error(err),
			)
		} else {
			logger.Log.Info(
				"Request completed",
				zap.Int("status", statusCode),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("ip", ip),
				zap.Duration("duration", duration),
			)
		}
		return err
	}
}

// RequestID adds a unique request ID to the context and response headers
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get(fiber.HeaderXRequestID)
		if requestID == "" {
			requestID = fmt.Sprintf("%d-%s", time.Now().UnixNano(), c.IP()) // Simple ID generation
		}
		c.Locals("requestid", requestID) // Store in Fiber's locals
		c.Set(fiber.HeaderXRequestID, requestID)
		return c.Next()
	}
}
