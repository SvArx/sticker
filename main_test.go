// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Testing Helpers
func runRequest(req *http.Request) (echo.Context, *httptest.ResponseRecorder){
	e := setUpUrlHandlers()

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return c, rec;
}

// Test Endpoints

func TestHelloEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/HelloWorld", nil)
	c, rec := runRequest(req)

	if assert.NoError(t, c.String(http.StatusOK, "Hello World")) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello World", rec.Body.String())
	}
}

func TestIndex(t *testing.T){
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	_, rec := runRequest(req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

// Test Module
