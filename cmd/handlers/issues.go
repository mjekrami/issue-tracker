package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllIssues(c echo.Context) error {
	return c.String(http.StatusNotFound, "To be implemented")
}
