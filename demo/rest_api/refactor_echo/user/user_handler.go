package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	Service UserServicer
}

func NewUserHandler(e *echo.Echo, service UserServicer) {
	h := UserHandler{Service: service}
	e.GET("/user", h.GetUser)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := h.Service.GetUser(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	// TODO with errors
	return http.StatusInternalServerError
}
