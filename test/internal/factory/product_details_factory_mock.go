package factory

import "github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"

var MockProductItem = domain.ProductItem{
	ProductItemID: "1",
	SellerID:      "S001",
	ProductID:     "P001",
	Price:         100.0,
	PreviousPrice: 120.0,
	Stock:         10,
	Color:         "Preto",
	Photos:        []string{"photo1.jpg", "photo2.jpg"},
	BestSeller:    true,
	Product:       MockProduct,
}

var MockPaymentMethods = []domain.PaymentMethod{
	{
		PaymentMethodID: "visa",
		Icon:            "visa.png",
		Type:            "credit",
	},
	{
		PaymentMethodID: "master",
		Icon:            "mastercard.png",
		Type:            "debit",
	},
}

var MockProduct = domain.Product{
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
}
