package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/handlers"
)

func Auth(e *echo.Echo) {
	g := e.Group("/auth")
	g.Add("POST", "/login", handlers.LoginHandler)
}
