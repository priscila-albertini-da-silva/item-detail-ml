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

func createTempPaymentMethodsFile(t *testing.T, data []domain.PaymentMethod) (string, func()) {
	dir, err := ioutil.TempDir("", "payment_methods_test")
	assert.NoError(t, err)

	filePath := filepath.Join(dir, "payment_methods.json")
	bytes, err := json.Marshal(data)
	assert.NoError(t, err)

	err = ioutil.WriteFile(filePath, bytes, 0644)
	assert.NoError(t, err)

	cleanup := func() {
		os.RemoveAll(dir)
	}
	return dir, cleanup
}

func TestPaymentMethodRepository_GetAll_Success(t *testing.T) {
	dir, cleanup := createTempPaymentMethodsFile(t, MockPaymentMethods)
	defer cleanup()

	repo := repository.NewPaymentMethodRepository(dir)
	result, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, MockPaymentMethods, result)
}

func TestPaymentMethodRepository_GetAll_FileNotFound(t *testing.T) {
	dir, err := ioutil.TempDir("", "payment_methods_test")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	repo := repository.NewPaymentMethodRepository(dir)
	_, err = repo.GetAll()
	assert.Error(t, err)
}

func TestPaymentMethodRepository_GetAll_InvalidJSON(t *testing.T) {
	dir, err := ioutil.TempDir("", "payment_methods_test")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	filePath := filepath.Join(dir, "payment_methods.json")
	err = ioutil.WriteFile(filePath, []byte("invalid json"), 0644)
	assert.NoError(t, err)

	repo := repository.NewPaymentMethodRepository(dir)
	_, err = repo.GetAll()
	assert.Error(t, err)
}
