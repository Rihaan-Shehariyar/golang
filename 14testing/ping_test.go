package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPing(t *testing.t) {
	// Gin test mode (important)
	gin.SetMode(gin.TestMode)

	// Fake response writer
	w := httptest.NewRecorder()

	// Router
	r := gin.Default()
	r.GET("/ping", ping)

	// Fake HTTP request
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Serve request
	r.ServeHTTP(w, req)

	// Assert status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}
