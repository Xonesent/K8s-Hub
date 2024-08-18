package middleware

import "github.com/gofiber/fiber/v2"

type MDWManager interface {
	NoMiddleware() fiber.Handler
}
