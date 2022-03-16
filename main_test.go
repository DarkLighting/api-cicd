package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func appInit() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", HomeHandler)
	return router
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHomeHandlerStatus(t *testing.T) {
	router := appInit()
	w := performRequest(router, "GET", "/")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHomeHandlerBody(t *testing.T) {
	router := appInit()
	w := performRequest(router, "GET", "/")
	assert.Equal(t, "Hello World!\n", w.Body.String())
}

func TestHomeHandlerMethods(t *testing.T) {
	router := appInit()
	methods := []string{"POST", "PUT", "DELETE", "OPTIONS", "TRACE", "PATCH"}
	for _, m := range methods {
		w := performRequest(router, m, "/")
		assert.Equal(t, http.StatusOK, w.Code)
	}
}
