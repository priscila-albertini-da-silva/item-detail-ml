package integration

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/factory"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/repository"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func setupIntegrationRouter(dataPath string) *gin.Engine {
	productRepo := repository.NewProductRepository(dataPath)
	paymentMethodRepo := repository.NewPaymentMethodRepository(dataPath)

	paymentMethodUC := usecase.NewPaymentMethodUseCase(paymentMethodRepo)
	responseFactory := factory.NewProductDetailResponseFactory()

	uc := usecase.NewProductDetailUseCase(productRepo, paymentMethodUC, responseFactory)
	h := handler.NewProductHandler(uc)

	r := gin.Default()
	r.GET("/products/:id", h.GetProductDetails)
	r.GET("/products", h.GetProductDetails)
	return r
}

func TestIntegration_GetProductDetails_Success(t *testing.T) {
	dir := t.TempDir()
	productsJSON := `[{"ProductItemID": "1", "ProductID": "P001", "Product": { "ProductID": "P001", "Name": "Produto Teste" }}]`
	err := os.WriteFile(filepath.Join(dir, "products.json"), []byte(productsJSON), 0644)
	assert.NoError(t, err)

	router := setupIntegrationRouter(dir)
	req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Produto Teste")
}

func TestIntegration_GetProductDetails_NotFound(t *testing.T) {
	dir := t.TempDir()
	productsJSON := `[{"ProductItemID": "PI001", "ProductID": "P001", "Product": { "ProductID": "P001", "Name": "Produto Teste" }}]`
	_ = os.WriteFile(filepath.Join(dir, "products.json"), []byte(productsJSON), 0644)

	router := setupIntegrationRouter(dir)
	req, _ := http.NewRequest(http.MethodGet, "/products/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Product not found")
}

func TestIntegration_GetProductDetails_MissingID(t *testing.T) {
	dir := t.TempDir()
	router := setupIntegrationRouter(dir)
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Product Item ID is required")
}

func TestIntegration_GetProductDetails_InternalError(t *testing.T) {
	dir := t.TempDir()
	invalidJSON := `{{{{{{`
	err := os.WriteFile(filepath.Join(dir, "products.json"), []byte(invalidJSON), 0644)
	assert.NoError(t, err)

	router := setupIntegrationRouter(dir)
	req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Internal")
}
