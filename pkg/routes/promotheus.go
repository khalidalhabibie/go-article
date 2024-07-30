package routes

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PromotrheusRoute func for display metrics
func PromotheusrRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/metrics")

	// Routes for GET method:
	route.Get("*", adaptor.HTTPHandler(promhttp.Handler()))
}
