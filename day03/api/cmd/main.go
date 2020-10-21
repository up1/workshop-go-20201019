package main

import (
	"demo/user"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create server with Echo
	e := gin.Default()
	pprof.Register(e)

	// Initial dependencies
	userService := user.NewUserService()
	user.NewUserHandler(e, userService)

	// Start server
	e.Run(":8080")
}
