package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/handlers"
)

func UsersRoute(e *echo.Echo) {
	userRoutes := e.Group("/users")
	userRoutes.Add("GET", "/all", handlers.GetAllUsers).Name = "get-all-users"
	userRoutes.Add("GET", "/:id", handlers.GetUserByID).Name = "get-user-by-id"
	userRoutes.Add("POST", "/create-user", handlers.CreateUser).Name = "create-user"
}
