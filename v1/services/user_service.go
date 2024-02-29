package services

import "github.com/labstack/echo/v4"

type (
	UserService struct{}
)

func (u UserService) UserByID(id string) (data string, err *echo.HTTPError) {
	return "Hello World", nil
}
