package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/health", handler.HealthCheck)

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"status":"ok"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
