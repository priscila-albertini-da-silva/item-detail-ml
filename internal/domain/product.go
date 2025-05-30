package domain

type Product struct {
	ProductID         string
	Name              string
	Description       string
	Category          Category
	Brand             string
	RelatedProducts   []RelatedProduct
	TechnicalFeatures map[string]interface{}
}

type RelatedProduct struct {
	ProductItemID string
	Name          string
	Price         float64
	Promotion     string
	Photo         string
}
