package usecase

import (
	"testing"

	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler/delivery"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/usecase"
	"github.com/priscila-albertini-da-silva/item-detail-ml/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductUseCase_GetAll(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	mockRepo.On("GetAll").Return([]domain.ProductItem{{ProductItemID: "1"}}, nil).Once()

	uc := usecase.NewProductUseCase(mockRepo, nil, nil)
	got, err := uc.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, []domain.ProductItem{{ProductItemID: "1"}}, got)
	mockRepo.AssertExpectations(t)
}

func TestProductUseCase_GetProduct(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	mockRepo.On("GetAll").Return([]domain.ProductItem{{ProductItemID: "1"}}, nil).Once()

	uc := usecase.NewProductUseCase(mockRepo, nil, nil)
	got, err := uc.GetProduct("1")
	assert.NoError(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, "1", got.ProductItemID)
	mockRepo.AssertExpectations(t)
}

func TestProductUseCase_GetProductDetails(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	mockPaymentUC := new(mocks.PaymentMethodUseCase)

	mockRepo.On("GetAll").Return([]domain.ProductItem{{ProductItemID: "1"}}, nil).Once()
	mockPaymentUC.On("GetAll").Return([]domain.PaymentMethod{{PaymentMethodID: "pm1"}}, nil).Once()

	mockFactory := new(mocks.ItemProductResponseFactory)
	mockFactory.On("ToItemProductResponse", mock.Anything, mock.Anything).
		Return(&delivery.ItemProductResponse{ProductItemID: "1"})
	uc := usecase.NewProductUseCase(mockRepo, mockPaymentUC, mockFactory)

	got, err := uc.GetProductDetails("1")
	assert.NoError(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, "1", got.ProductItemID)
	mockRepo.AssertExpectations(t)
	mockPaymentUC.AssertExpectations(t)
}
