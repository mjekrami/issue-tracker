package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/database"
)

func GetAllIssues(c echo.Context) error {
	res, err := database.GetAllIssues()
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could get the data from database"})
	}
	return c.JSON(http.StatusOK, res)
}

func GetIssue(c echo.Context) error {
	name := c.Param("issueName")
	res, err := database.GetIssue(name)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not get the data from database"})
	}
	return c.JSON(http.StatusOK, res)
}
