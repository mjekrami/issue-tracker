package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/handlers"
)

func IssuesRoute(e *echo.Echo) {
	issueRoutes := e.Group("/issues")
	issueRoutes.Add("GET", "/all", handlers.GetAllIssues).Name = "get-all-issues"
	issueRoutes.Add("GET", "/:issueName", handlers.GetIssue).Name = "get-issues-by-name"
}
