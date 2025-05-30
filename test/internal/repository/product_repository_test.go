package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/repository"
	"github.com/stretchr/testify/assert"
)

func createTempProductsFile(t *testing.T, products []domain.ProductItem) (string, func()) {
	t.Helper()
	dir, err := ioutil.TempDir("", "productrepo")
	assert.NoError(t, err)

	filePath := filepath.Join(dir, "products.json")
	data, err := json.Marshal(products)
	assert.NoError(t, err)

	err = ioutil.WriteFile(filePath, data, 0644)
	assert.NoError(t, err)

	cleanup := func() {
		os.RemoveAll(dir)
	}
	return dir, cleanup
}

func TestGetAll_ReturnsProducts(t *testing.T) {
	dir, cleanup := createTempProductsFile(t, MockProductItems)
	defer cleanup()

	repo := repository.NewProductRepository(dir)
	products, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, MockProductItems, products)
}

func TestGetAll_FileNotFound(t *testing.T) {
	dir, err := ioutil.TempDir("", "productrepo")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	repo := repository.NewProductRepository(dir)
	_, err = repo.GetAll()
	assert.Error(t, err)
}

func TestGetAll_InvalidJSON(t *testing.T) {
	dir, err := ioutil.TempDir("", "productrepo")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	filePath := filepath.Join(dir, "products.json")
	err = ioutil.WriteFile(filePath, []byte("{invalid json"), 0644)
	assert.NoError(t, err)

	repo := repository.NewProductRepository(dir)
	_, err = repo.GetAll()
	assert.Error(t, err)
}
