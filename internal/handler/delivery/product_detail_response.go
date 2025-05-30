package delivery

type ProductDetailResponse struct {
	ProductItemID     string
	ProductID         string
	Name              string
	Description       string
	Price             float64
	PreviousPrice     float64
	Stock             int
	BestSeller        bool
	Brand             string
	Photos            []string
	Category          ProductCategoryResponse
	Colors            []ProductColorResponse
	Promotions        []ProductPromotionResponse
	RelatedProducts   []RelatedProductsResponse
	Reviews           []ProductReviewResponse
	Seller            SellerResponse
	TechnicalFeatures map[string]interface{}
	PaymentMethods    []PaymentMethodResponse
}

type ProductColorResponse struct {
	ProductItemID string
	Color         string
	Photo         string
}

type ProductCategoryResponse struct {
	CategoryID     string
	Description    string
	ParentCategory *ProductCategoryResponse
}

type ProductPromotionResponse struct {
	ProductID        string
	Description      string
	DiscountPct      float64
	Installments     int
	InstallmentValue float64
	CardType         string
	ExtraDiscountPct float64
}

type ProductReviewResponse struct {
	ReviewID string
	Customer string
	Rate     int
	Review   string
}

type RelatedProductsResponse struct {
	ProductItemID string
	Name          string
	Photo         string
	Price         float64
	Promotion     string
}

type SellerResponse struct {
	SellerID     string
	Name         string
	OfficialSeal bool
	Rating       float64
}

type PaymentMethodResponse struct {
	PaymentMethodID string
	Icon            string
	Type            string
}
