package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	Service UserServicer
}

func NewUserHandler(e *gin.Engine, service UserServicer) {
	h := UserHandler{Service: service}
	e.GET("/user", h.GetUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := h.Service.GetUser(ctx)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	c.JSON(http.StatusOK, user)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	// TODO with errors
	return http.StatusInternalServerError
}
