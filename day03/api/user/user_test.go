package user_test

import (
	"context"
	"demo/user"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type stubUserService struct {
	err error
}

func (s stubUserService) GetUser(c context.Context) (user.User, error) {
	user := user.User{Firstname: "user test"}
	return user, s.err
}

func TestSuccessWithGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request, _ = http.NewRequest(http.MethodGet, "/user", nil)
	handler := user.UserHandler{
		Service: stubUserService{},
	}
	handler.GetUser(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"firstname":"user test","lastname":"","title":""}`, rec.Body.String())
}

func TestFailWithGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request, _ = http.NewRequest(http.MethodGet, "/user", nil)
	handler := user.UserHandler{
		Service: stubUserService{err: fmt.Errorf("Error")},
	}
	handler.GetUser(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
