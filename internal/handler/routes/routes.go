package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/priscila-albertini-da-silva/item-detail-ml/docs"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine, productHandler *handler.ProductDetailHandler) {
	r.GET("/products/:id", productHandler.GetProductDetails)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", handler.HealthCheck)
}
