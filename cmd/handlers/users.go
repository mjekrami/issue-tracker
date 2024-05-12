package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/database"
	"github.com/mjekrami/issue-tracker/cmd/models"
)

var (
	UserNotFoundErr     = userResponseErr{"user not found"}
	UserCannotCreateErr = userResponseErr{"cannot create user"}
)

type userResponseErr struct {
	Reason string `json:"reason" required:"true"`
}

func HandleGetAllUsers(c echo.Context) error {
	users, err := database.GetAllUsers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
func HandleGetUserByID(c echo.Context) error {
	Id := c.Param("id")
	user, err := database.GetUserById(Id)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, UserNotFoundErr)
	}
	return c.JSON(http.StatusOK, user)
}
func HandleGetUserByName(c echo.Context) error {
	name := c.Param("name")
	user, err := database.GetUserByName(name)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, UserNotFoundErr)
	}
	return c.JSON(http.StatusOK, user)
}

func HandleCreateUser(c echo.Context) error {
	u := new(models.UserCreate)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.CreateUser(u.Name, u.LastName, u.Privilege); err != nil {
		return c.JSON(http.StatusBadRequest, UserCannotCreateErr)
	}
	userResponse := &models.UserResponse{
		UserCreate: u,
		JoinDate:   time.Now().String(),
	}
	return c.JSON(http.StatusOK, userResponse)
}
