package v1

import (
	"echo-template/v1/routes"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	v1 := e.Group("/v1")

	routes.UserGroup(v1)
}
