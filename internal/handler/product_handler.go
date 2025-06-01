package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler/delivery"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/usecase"
)

type ProductDetailHandler struct {
	usecase usecase.ProductDetailUseCase
}

func NewProductHandler(uc usecase.ProductDetailUseCase) *ProductDetailHandler {
	return &ProductDetailHandler{usecase: uc}
}

// GetProductDetails godoc
// @Summary      Detalhes do produto
// @Description  Retorna detalhes do produto por ID do item
// @Tags         products
// @Param        id   path      string  true  "Product Item ID"
// @Success      200  {object}  delivery.ProductDetailResponse
// @Failure      400  {object}  delivery.ErrorResponse "ID ausente ou inválido"
// @Failure      404  {object}  delivery.ErrorResponse "Produto não encontrado"
// @Failure      500  {object}  delivery.ErrorResponse "Erro interno"
// @Router       /products/{id} [get]
func (h *ProductDetailHandler) GetProductDetails(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, delivery.ErrorResponse{Error: "Product Item ID is required"})
		return
	}

	productDetails, err := h.usecase.GetProductDetails(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, delivery.ErrorResponse{Error: "Internal server error: " + err.Error()})
		return
	}

	if productDetails == nil {
		c.JSON(http.StatusNotFound, delivery.ErrorResponse{Error: "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": productDetails})
}
