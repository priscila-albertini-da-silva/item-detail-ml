package usecase

import (
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/factory"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler/delivery"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/repository"
	"github.com/sirupsen/logrus"
)

type ProductDetailUseCase interface {
	GetProductItem(productID string) (*domain.ProductItem, error)
	GetProductDetails(productID string) (*delivery.ProductDetailResponse, error)
}

type productDetailUseCase struct {
	productItemRepository         repository.ProductItemRepository
	paymentMethodUseCase          PaymentMethodUseCase
	productDetailsResponseFactory factory.ProductDetailResponseFactory
}

func NewProductDetailUseCase(repo repository.ProductItemRepository, paymentMethodUseCase PaymentMethodUseCase, factory factory.ProductDetailResponseFactory) *productDetailUseCase {
	return &productDetailUseCase{productItemRepository: repo, paymentMethodUseCase: paymentMethodUseCase, productDetailsResponseFactory: factory}
}

func (uc *productDetailUseCase) GetProductDetails(productItemID string) (*delivery.ProductDetailResponse, error) {
	logrus.Infof("Fetching details for product item ID: %s", productItemID)

	product, err := uc.GetProductItem(productItemID)
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
	return uc.productDetailsResponseFactory.ToProductDetailResponse(*product, paymentMethods), nil
}

func (uc *productDetailUseCase) GetProductItem(productItemID string) (*domain.ProductItem, error) {
	logrus.Infof("Fetching product with item ID: %s", productItemID)

	products, err := uc.productItemRepository.GetAll()
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
