package middleware

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	adaptor "github.com/gofiber/adaptor/v2"

	"github.com/gofiber/helmet/v2"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {

	metricsMiddleware := NewMetricsMiddleware()
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		logger.New(),
		//
		helmet.New(),

		// csrf.New(), // add CSRF middleware

		adaptor.HTTPMiddleware(metricsMiddleware.Metrics),

		// adaptor.HTTPHandlerFunc(NewMetricsMiddleware().Metrics()),
	)
}
