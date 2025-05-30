package usecase

import "github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler/delivery"

var MockProductItem = delivery.ItemProductResponse{
	ProductItemID: "P001",
	ProductID:     "1",
	Name:          "Produto Mock",
	Description:   "Descrição do produto mock",
	Brand:         "Marca Mock",
	Price:         100.0,
	PreviousPrice: 120.0,
	Stock:         10,
	Photos:        []string{"mock1.jpg", "mock2.jpg"},
	BestSeller:    true,
}
