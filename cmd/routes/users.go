package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/handlers"
)

func UsersRoute(e *echo.Echo) {
	userRoutes := e.Group("/users")
	// GET
	userRoutes.Add("GET", "/all", handlers.HandleGetAllUsers).Name = "get-all-users"
	userRoutes.Add("GET", "/:id", handlers.HandleGetUserByID).Name = "get-user-by-id"
	userRoutes.Add("GET", "/:name", handlers.HandleGetUserByName).Name = "get-user-by-name"

	// POST
	userRoutes.Add("POST", "/create", handlers.HandleCreateUser).Name = "create-user"
}
