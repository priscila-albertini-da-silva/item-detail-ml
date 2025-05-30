package usecase

import (
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/factory"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler/delivery"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/repository"
	"github.com/sirupsen/logrus"
)

type ProductUseCase interface {
	GetAll() ([]domain.ProductItem, error)
	GetProduct(productID string) (*domain.ProductItem, error)
	GetProductDetails(productID string) (*delivery.ItemProductResponse, error)
}

type productUseCase struct {
	repo                 repository.ProductRepository
	paymentMethodUseCase PaymentMethodUseCase
	factory              factory.ItemProductResponseFactory
}

func NewProductUseCase(repo repository.ProductRepository, paymentMethodUseCase PaymentMethodUseCase, factory factory.ItemProductResponseFactory) *productUseCase {
	return &productUseCase{repo: repo, paymentMethodUseCase: paymentMethodUseCase, factory: factory}
}

func (uc *productUseCase) GetProductDetails(productItemID string) (*delivery.ItemProductResponse, error) {
	logrus.Infof("Fetching details for product item ID: %s", productItemID)

	product, err := uc.GetProduct(productItemID)
	if err != nil {
		logrus.Errorf("Error fetching product details: %v", err)
		return nil, err
	}
	if product == nil {
		logrus.Warnf("No product found with item ID: %s", productItemID)
		return nil, nil
	}

	paymentMethods, err := uc.paymentMethodUseCase.GetAll()
	if err != nil {
		logrus.Errorf("Error fetching payment methods: %v", err)
		return nil, err
	}

	logrus.Infof("Successfully fetched product details for item ID: %s", productItemID)
	return uc.factory.ToItemProductResponse(*product, paymentMethods), nil
}

func (uc *productUseCase) GetProduct(productItemID string) (*domain.ProductItem, error) {
	logrus.Infof("Fetching product with item ID: %s", productItemID)

	products, err := uc.repo.GetAll()
	if err != nil {
		logrus.Errorf("Error fetching products: %v", err)
		return nil, err
	}
	if len(products) == 0 {
		logrus.Warn("No products found in the repository")
		return nil, nil
	}
	for _, p := range products {
		if p.ProductItemID == productItemID {
			logrus.Infof("Found product with item ID: %s", productItemID)
			return &p, nil
		}
	}

	logrus.Warnf("No product found with item ID: %s", productItemID)
	return nil, nil
}

func (uc *productUseCase) GetAll() ([]domain.ProductItem, error) {
	logrus.Info("Fetching all products")

	products, err := uc.repo.GetAll()
	if err != nil {
		logrus.Errorf("Error fetching all products: %v", err)
		return nil, err
	}

	if len(products) == 0 {
		logrus.Warn("No products found in the repository")
		return nil, nil
	}

	logrus.Infof("Successfully fetched %d products", len(products))
	return products, nil
}
