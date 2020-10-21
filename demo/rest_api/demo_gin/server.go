package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", getUser)
	r.Run(":8080")
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Title     string `json:"title"`
}

func getUser(c *gin.Context) {
	u := &User{
		Firstname: "F01",
		Lastname:  "L01",
	}
	c.JSON(http.StatusOK, u)
}
