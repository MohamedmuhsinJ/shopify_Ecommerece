package controllers

import (
	"bytes"
	"encoding/json"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestEditProduct(t *testing.T) {
	r := gin.Default()
	r.POST("/editproduct/:id", func(c *gin.Context) {

	})
	// var product Editproduct
	price := 100000
	prod, _ := json.Marshal(price)
	req, err := http.NewRequest("POST", "/editproduct/2", bytes.NewBuffer(prod))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}
