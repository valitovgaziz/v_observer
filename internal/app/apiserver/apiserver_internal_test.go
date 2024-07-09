package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	s.handleHello().ServeHTTP(res, req)
	assert.Equal(t, res.Body.String(), "Hello")
}
