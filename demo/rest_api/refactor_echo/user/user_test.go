package user_test

import (
	"context"
	"demo/user"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubUserService struct {
	err error
}

func (s stubUserService) GetUser(c context.Context) (user.User, error) {
	user := user.User{Firstname: "user test"}
	return user, s.err
}

func TestSuccessWithGetUser(t *testing.T) {

	e := echo.New()
	req, err := http.NewRequest(http.MethodGet, "/user", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := user.UserHandler{
		Service: stubUserService{},
	}
	err = handler.GetUser(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"firstname":"user test","lastname":"","title":""}`, rec.Body.String())
}

func TestFailWithGetUser(t *testing.T) {

	e := echo.New()
	req, err := http.NewRequest(http.MethodGet, "/userd", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := user.UserHandler{
		Service: stubUserService{err: fmt.Errorf("Error")},
	}
	err = handler.GetUser(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
