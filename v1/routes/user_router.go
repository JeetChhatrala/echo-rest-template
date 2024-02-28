package routes

import (
	"echo-template/v1/controllers"

	"github.com/labstack/echo/v4"
)

func UserGroup(parent *echo.Group) {
	userGroup := parent.Group("/users")
	userController := controllers.UserController{}

	userGroup.GET("/", userController.GetUserByID)
}
