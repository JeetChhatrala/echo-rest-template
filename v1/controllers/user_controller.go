package controllers

import (
	"echo-template/v1/models"
	"echo-template/v1/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	UserController struct {
		services services.UserService
	}
)

/*This controller is used to retrive user by userID*/
func (u *UserController) GetUserByID(c echo.Context) error {

	user := new(models.User)

	// Binding request data to object
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Call service and return response accordingly
	if data, err := u.services.UserByID(user.ID); err != nil {
		return err
	} else {
		return c.String(http.StatusOK, data)
	}
}
