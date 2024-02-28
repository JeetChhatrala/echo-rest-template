package controllers

import (
	"echo-template/v1/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	UserController struct {
		services services.UserService
	}
)

func (u *UserController) GetUserByID(c echo.Context) error {
	if res, err := u.services.UserByID("ID"); err != nil {
		return c.String(http.StatusExpectationFailed, res)
	} else {
		return c.String(http.StatusOK, res)
	}
}
