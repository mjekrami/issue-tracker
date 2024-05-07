package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/dtos"
	"net/http"
)

func GetUserByID(c echo.Context) error {
	Id := c.Param("id")
	res := map[string]string{"user": Id}
	return c.JSON(http.StatusOK, res)
}

func CreateUser(c echo.Context) error {
	u := new(dtos.User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	user := dtos.UserDTO{
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
	return c.JSON(http.StatusOK, user)

}
