package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/factory"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler/routes"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/repository"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/usecase"
	"github.com/sirupsen/logrus"
)

const (
	cmdFolder      = "cmd"
	mockDataFolder = "mock_data"
	serverPort     = ":8080"
)

func main() {
	logrus.Info("Starting Item Detail ML Server...")

	cwd, _ := os.Getwd()
	if filepath.Base(cwd) == cmdFolder {
		cwd = filepath.Dir(cwd)
	}
	mockDataPath := filepath.Join(cwd, mockDataFolder)

	productRepo := repository.NewProductRepository(mockDataPath)
	paymentMethodRepo := repository.NewPaymentMethodRepository(mockDataPath)
	paymentMethodUseCase := usecase.NewPaymentMethodUseCase(paymentMethodRepo)
	productDetailsFactory := factory.NewProductDetailResponseFactory()
	ProductDetailUseCase := usecase.NewProductDetailUseCase(productRepo, paymentMethodUseCase, productDetailsFactory)
	productHandler := handler.NewProductHandler(ProductDetailUseCase)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	routes.RegisterRoutes(r, productHandler)
	r.Run(serverPort)
}
