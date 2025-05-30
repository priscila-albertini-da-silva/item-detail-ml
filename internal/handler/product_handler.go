package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/usecase"
)

type ProductDetailHandler struct {
	usecase usecase.ProductDetailUseCase
}

func NewProductHandler(uc usecase.ProductDetailUseCase) *ProductDetailHandler {
	return &ProductDetailHandler{usecase: uc}
}

func (h *ProductDetailHandler) GetProductDetails(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	productDetails, err := h.usecase.GetProductDetails(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": productDetails})
}
