package main

import (
	v1 "echo-template/v1"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// API logging middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "status=${status}, method=${method}, latency=${latency_human}, uri=${uri}, error_if_any=${error}\n",
	}))

	// Initialise v1 routes
	v1.Routes(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
