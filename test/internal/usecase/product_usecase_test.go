package usecase

import (
	"errors"
	"testing"

	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler/delivery"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/usecase"
	"github.com/priscila-albertini-da-silva/item-detail-ml/test/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductDetailUseCase_GetAll(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	mockRepo.On("GetAll").Return([]domain.ProductItem{{ProductItemID: "1"}}, nil).Once()

	uc := usecase.NewProductDetailUseCase(mockRepo, nil, nil)
	got, err := uc.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, []domain.ProductItem{{ProductItemID: "1"}}, got)
	mockRepo.AssertExpectations(t)
}

func TestProductDetailUseCase_GetProduct(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	mockRepo.On("GetAll").Return([]domain.ProductItem{{ProductItemID: "1"}}, nil).Once()

	uc := usecase.NewProductDetailUseCase(mockRepo, nil, nil)
	got, err := uc.GetProductItem("1")
	assert.NoError(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, "1", got.ProductItemID)
	mockRepo.AssertExpectations(t)
}

func TestProductDetailUseCase_GetProductDetails(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	mockPaymentUC := new(mocks.PaymentMethodUseCase)

	mockRepo.On("GetAll").Return([]domain.ProductItem{{ProductItemID: "1"}}, nil).Once()
	mockPaymentUC.On("GetAll").Return([]domain.PaymentMethod{{PaymentMethodID: "pm1"}}, nil).Once()

	mockFactory := new(mocks.ProductDetailResponseFactory)
	mockFactory.On("ToProductDetailResponse", mock.Anything, mock.Anything).
		Return(&delivery.ProductDetailResponse{ProductItemID: "1"})
	uc := usecase.NewProductDetailUseCase(mockRepo, mockPaymentUC, mockFactory)

	got, err := uc.GetProductDetails("1")
	assert.NoError(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, "1", got.ProductItemID)
	mockRepo.AssertExpectations(t)
	mockPaymentUC.AssertExpectations(t)
}

func TestGetProductDetails_ErrorOnGetProduct(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	mockPaymentUC := new(mocks.PaymentMethodUseCase)
	mockFactory := new(mocks.ProductDetailResponseFactory)

	uc := usecase.NewProductDetailUseCase(mockRepo, mockPaymentUC, mockFactory)

	mockRepo.On("GetAll").Return(nil, errors.New("repo error"))

	resp, err := uc.GetProductDetails("1")
	assert.Nil(t, resp)
	assert.EqualError(t, err, "repo error")
}

func TestGetProductDetails_ProductNil(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	mockPaymentUC := new(mocks.PaymentMethodUseCase)
	mockFactory := new(mocks.ProductDetailResponseFactory)

	uc := usecase.NewProductDetailUseCase(mockRepo, mockPaymentUC, mockFactory)

	mockRepo.On("GetAll").Return(nil, nil)

	resp, err := uc.GetProductDetails("1")
	assert.Nil(t, resp)
	assert.NoError(t, err)
}

func TestGetProductDetails_ErrorOnGetPaymentMethods(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	mockPaymentUC := new(mocks.PaymentMethodUseCase)
	mockFactory := new(mocks.ProductDetailResponseFactory)

	uc := usecase.NewProductDetailUseCase(mockRepo, mockPaymentUC, mockFactory)

	mockRepo.On("GetAll").Return([]domain.ProductItem{{ProductItemID: "1"}}, nil).Once()

	mockPaymentUC.On("GetAll").Return(nil, errors.New("payment error"))

	resp, err := uc.GetProductDetails("1")
	assert.Nil(t, resp)
	assert.EqualError(t, err, "payment error")
}

func TestGetProduct_NotFound(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)

	mockRepo.On("GetAll").Return([]domain.ProductItem{
		{ProductItemID: "A"},
		{ProductItemID: "B"},
	}, nil)

	uc := usecase.NewProductDetailUseCase(mockRepo, nil, nil)

	result, err := uc.GetProductItem("X")
	assert.Nil(t, result)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAll_ErrorOnRepo(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	mockRepo.On("GetAll").Return(nil, errors.New("repo error"))

	uc := usecase.NewProductDetailUseCase(mockRepo, nil, nil)

	result, err := uc.GetAll()
	assert.Nil(t, result)
	assert.EqualError(t, err, "repo error")
	mockRepo.AssertExpectations(t)
}

func TestGetAll_EmptyProducts(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	mockRepo.On("GetAll").Return([]domain.ProductItem{}, nil)

	uc := usecase.NewProductDetailUseCase(mockRepo, nil, nil)

	result, err := uc.GetAll()
	assert.Nil(t, result)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
