package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/database"
)

var (
	IssueNotFoundErr = issueNotFoundErr{"issue not found"}
)

type issueNotFoundErr struct {
	Reason string `json:"reason" required:"true"`
}

func HandleGetAllIssues(c echo.Context) error {
	res, err := database.GetAllIssues()
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}
	return c.JSON(http.StatusOK, res)
}

func HandleGetIssue(c echo.Context) error {
	name := c.Param("issueName")
	res, err := database.GetIssue(name)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}
	if res == nil {
		return c.JSON(http.StatusNotFound, IssueNotFoundErr)
	}
	return c.JSON(http.StatusOK, res)
}
