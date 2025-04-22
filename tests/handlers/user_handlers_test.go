package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetPath("/get-user?id=1")
	assert.Equal(t, http.StatusOK, resp.Code)
}
