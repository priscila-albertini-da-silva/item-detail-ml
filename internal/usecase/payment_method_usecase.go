package usecase

import (
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/repository"
	"github.com/sirupsen/logrus"
)

type PaymentMethodUseCase interface {
	GetAll() ([]domain.PaymentMethod, error)
}

type paymentMethodUseCase struct {
	repo repository.PaymentMethodRepository
}

func NewPaymentMethodUseCase(repo repository.PaymentMethodRepository) *paymentMethodUseCase {
	return &paymentMethodUseCase{repo: repo}
}

func (uc *paymentMethodUseCase) GetAll() ([]domain.PaymentMethod, error) {
	logrus.Info("Fetching all payment methods")

	paymentMethods, err := uc.repo.GetAll()
	if err != nil {
		logrus.Errorf("Error fetching payment methods: %v", err)
		return nil, err
	}

	if len(paymentMethods) == 0 {
		logrus.Warn("No payment methods found in the repository")
		return nil, nil
	}
	logrus.Infof("Successfully fetched %d payment methods", len(paymentMethods))
	return paymentMethods, nil
}
