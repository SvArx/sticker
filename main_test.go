// main_test.go
package main

import (
	"reflect"
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


// Testcases

func TestSerilaizeAndDeserializeCart(t *testing.T) {
	num := 1;

	if !reflect.DeepEqual(num, 1) {
		t.Errorf("Example test.")
	}
}

func TestHelloEndpoint(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler function directly
	if assert.NoError(t, helloHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World!", rec.Body.String())
	}
}


// Pluming for tests
func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

