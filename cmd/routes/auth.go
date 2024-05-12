package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/handlers"
)

func AuthRoute(e *echo.Echo) {
	g := e.Group("/auth")
	g.Add("POST", "/login", handlers.HandleLogin)
}
