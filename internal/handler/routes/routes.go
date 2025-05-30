package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler"
)

func RegisterRoutes(r *gin.Engine, productHandler *handler.ProductDetailHandler) {
	r.GET("/product/:id", productHandler.GetProductDetails)
}
