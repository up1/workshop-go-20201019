package main

import (
	"demo/user"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create server with Echo
	e := echo.New()

	// Initial dependencies
	userService := user.NewUserService()
	user.NewUserHandler(e, userService)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
