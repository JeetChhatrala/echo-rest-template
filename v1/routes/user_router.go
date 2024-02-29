package routes

import (
	"echo-template/v1/controllers"
	"echo-template/v1/middlewares"

	"github.com/labstack/echo/v4"
)

func UserGroup(parent *echo.Group) {
	userGroup := parent.Group("/users")

	uc := controllers.UserController{}
	um := middlewares.UserMiddleware{}

	userGroup.GET("/:id", uc.GetUserByID, um.ValidateToken)
}
