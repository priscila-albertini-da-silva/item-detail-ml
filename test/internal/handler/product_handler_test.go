package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler/delivery"
	"github.com/priscila-albertini-da-silva/item-detail-ml/test/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func setupRouter(handler *handler.ProductHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/products/:id", handler.GetProductDetails)
	return r
}

func TestGetProductDetails_MissingID(t *testing.T) {
	mockUC := new(mocks.ProductUseCase)
	handler := handler.NewProductHandler(mockUC)
	router := gin.Default()
	router.GET("/products/", handler.GetProductDetails)

	req, _ := http.NewRequest(http.MethodGet, "/products/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Product ID is required")
}

func TestGetProductDetails_UseCaseError(t *testing.T) {
	mockUC := new(mocks.ProductUseCase)
	mockUC.On("GetProductDetails", "123").Return((*delivery.ItemProductResponse)(nil), errors.New("some error"))

	handler := handler.NewProductHandler(mockUC)
	router := setupRouter(handler)

	req, _ := http.NewRequest(http.MethodGet, "/products/123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "some error")
	mockUC.AssertExpectations(t)
}

func TestGetProductDetails_Success(t *testing.T) {
	expected := &delivery.ItemProductResponse{
		ProductID: "123",
		Name:      "Test Product",
	}
	mockUC := new(mocks.ProductUseCase)
	mockUC.On("GetProductDetails", "123").Return(expected, nil)

	handler := handler.NewProductHandler(mockUC)
	router := setupRouter(handler)

	req, _ := http.NewRequest(http.MethodGet, "/products/123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"ProductID":"123"`)
	assert.Contains(t, w.Body.String(), `"Name":"Test Product"`)
	mockUC.AssertExpectations(t)
}
