package domain

type ProductItem struct {
	ProductItemID string
	SellerID      string
	ProductID     string
	Seller        Seller
	Product       Product
	Price         float64
	PreviousPrice float64
	Stock         int
	Color         string
	Colors        []ProductColor
	Reviews       []Review
	Photos        []string
	Promotions    []Promotion
	BestSeller    bool
}

type ProductColor struct {
	ProductItemID string
	Color         string
	Photo         string
}
