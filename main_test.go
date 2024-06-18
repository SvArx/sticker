// main_test.go
package main

import (
	"reflect"
	"testing"
	"net/http"
	"net/http/httptest"

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
	e := setUpUrlHandlers()

	req := httptest.NewRequest(http.MethodGet, "/HelloWorld", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, c.String(http.StatusOK, "Hello World")) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello World", rec.Body.String())
	}
}
