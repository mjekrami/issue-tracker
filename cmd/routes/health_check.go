package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/handlers"
)

func HealthCheck(e *echo.Echo) {
	e.Add("GET", "/health", handlers.HealthCheckHandler).Name = "health-check"
}
