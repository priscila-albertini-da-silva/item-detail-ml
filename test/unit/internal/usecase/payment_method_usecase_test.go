package usecase

import (
	"errors"
	"testing"

	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/usecase"
	"github.com/priscila-albertini-da-silva/item-detail-ml/test/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPaymentMethodUseCase_GetAll_Success(t *testing.T) {
	mockRepo := new(mocks.PaymentMethodRepository)
	mockRepo.On("GetAll").Return(MockPaymentMethods, nil).Once()

	uc := usecase.NewPaymentMethodUseCase(mockRepo)

	result, err := uc.GetAll()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "pm1", result[0].PaymentMethodID)
	assert.Equal(t, "visa.png", result[0].Icon)
	assert.Equal(t, "credit", result[0].Type)
	assert.Equal(t, "pm2", result[1].PaymentMethodID)
	assert.Equal(t, "mastercard.png", result[1].Icon)
	assert.Equal(t, "debit", result[1].Type)
	mockRepo.AssertExpectations(t)
}

func TestPaymentMethodUseCase_GetAll_Empty(t *testing.T) {
	mockRepo := new(mocks.PaymentMethodRepository)
	mockRepo.On("GetAll").Return([]domain.PaymentMethod{}, nil).Once()

	uc := usecase.NewPaymentMethodUseCase(mockRepo)

	result, err := uc.GetAll()

	assert.NoError(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPaymentMethodUseCase_GetAll_Error(t *testing.T) {
	mockRepo := new(mocks.PaymentMethodRepository)
	mockRepo.On("GetAll").Return(nil, errors.New("repository error")).Once()

	uc := usecase.NewPaymentMethodUseCase(mockRepo)

	result, err := uc.GetAll()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
