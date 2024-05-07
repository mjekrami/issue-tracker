package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "mohammad" && password == "safePassword" {
		return c.JSON(http.StatusOK, map[string]string{"login": "sucessfull"})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{"login": "unsucessfull"})
}
