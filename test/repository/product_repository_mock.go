package repository

import "github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"

var MockProductItems = []domain.ProductItem{
	{
		ProductItemID: "1",
		SellerID:      "S001",
		ProductID:     "P001",
		Price:         100.0,
		PreviousPrice: 120.0,
		Stock:         10,
		Color:         "Preto",
		Photos:        []string{"photo1.jpg", "photo2.jpg"},
		BestSeller:    true,
		Product: domain.Product{
			ProductID:   "P001",
			Name:        "Smartphone X",
			Description: "Celular com câmera de 50MP",
			Brand:       "TechBrand",
			TechnicalFeatures: map[string]interface{}{
				"Camera":  "50MP",
				"Battery": "4000mAh",
			},
			Category: domain.Category{
				CategoryID:     "C001",
				Description:    "Celulares",
				ParentCategory: nil,
			},
		},
	},
	{
		ProductItemID: "2",
		SellerID:      "S002",
		ProductID:     "P002",
		Price:         2000.0,
		PreviousPrice: 2200.0,
		Stock:         5,
		Color:         "Vermelho",
		Photos:        []string{"bike1.jpg", "bike2.jpg"},
		BestSeller:    false,
		Product: domain.Product{
			ProductID:   "P002",
			Name:        "Mountain Bike",
			Description: "Bicicleta aro 29",
			Brand:       "BikeBrand",
			TechnicalFeatures: map[string]interface{}{
				"RimSize": "29",
				"Gears":   "24",
			},
			Category: domain.Category{
				CategoryID:     "C002",
				Description:    "Bicicletas",
				ParentCategory: nil,
			},
		},
	},
}
