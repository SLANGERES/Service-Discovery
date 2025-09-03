package Register

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SLANGERES/Service-Discovery/internal/models"
	"github.com/gin-gonic/gin"
)

func TestRegisterService_Success(t *testing.T) {
    // Mock the request body
    service := models.Service{
        Name: "test-service",
        Host: "localhost",
        Port: 8080,
    }
    jsonBody, _ := json.Marshal(service)
    req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")

    // Create a Gin router
    router := gin.Default()
    router.POST("/register", RegisterService)

    // Create a response recorder
    rr := httptest.NewRecorder()

    // Serve the request
    router.ServeHTTP(rr, req)

    // Assertions
    if rr.Code != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
    }

    expected := `{"message":"Service update sucessfully"}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
