package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLoginEndpoint(t *testing.T) {
	// Create a new Gin router instance
	r := gin.Default()
	var log struct {
		Email    string
		Password string
	}
	log.Email = "mm@gmail.com"
	log.Password = "1234"
	entry, _ := json.Marshal(log)
	r.POST("/login", func(c *gin.Context) {

	})

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(entry))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}
