package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/handlers"
)

func HealthCheckRoute(e *echo.Echo) {
	e.Add("GET", "/health", handlers.HandleHealthCheck).Name = "health-check"
}
