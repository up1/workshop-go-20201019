package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/user", getUser)
	e.Logger.Fatal(e.Start(":8080"))
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Title     string `json:"title"`
}

func getUser(c echo.Context) error {
	u := &User{
		Firstname: "F01",
		Lastname:  "L01",
	}
	return c.JSON(http.StatusOK, u)
}
