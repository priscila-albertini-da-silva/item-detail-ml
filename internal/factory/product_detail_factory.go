package factory

import (
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"
	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/handler/delivery"
)

type ProductDetailResponseFactory interface {
	ToProductDetailResponse(p domain.ProductItem, pm []domain.PaymentMethod) *delivery.ProductDetailResponse
}

type productDetailResponseFactory struct{}

func NewProductDetailResponseFactory() ProductDetailResponseFactory {
	return &productDetailResponseFactory{}
}

func (f *productDetailResponseFactory) ToProductDetailResponse(p domain.ProductItem, pm []domain.PaymentMethod) *delivery.ProductDetailResponse {
	return &delivery.ProductDetailResponse{
		ProductID:         p.ProductID,
		Price:             p.Price,
		PreviousPrice:     p.PreviousPrice,
		Stock:             p.Stock,
		Photos:            p.Photos,
		BestSeller:        p.BestSeller,
		Description:       p.Product.Description,
		Name:              p.Product.Name,
		Brand:             p.Product.Brand,
		TechnicalFeatures: p.Product.TechnicalFeatures,
		Category: delivery.ProductCategoryResponse{
			CategoryID:  p.Product.Category.CategoryID,
			Description: p.Product.Category.Description,
			ParentCategory: func() *delivery.ProductCategoryResponse {
				if p.Product.Category.ParentCategory == nil {
					return nil
				}
				return &delivery.ProductCategoryResponse{
					CategoryID:  p.Product.Category.ParentCategory.CategoryID,
					Description: p.Product.Category.ParentCategory.Description,
				}
			}(),
		},
		Reviews: func() []delivery.ProductReviewResponse {
			if len(p.Reviews) == 0 {
				return nil
			}
			reviews := make([]delivery.ProductReviewResponse, len(p.Reviews))
			for i, r := range p.Reviews {
				reviews[i] = delivery.ProductReviewResponse{
					ReviewID: r.ReviewID,
					Customer: r.Customer,
					Rate:     r.Rate,
					Review:   r.Review,
				}
			}
			return reviews
		}(),
		Colors: func() []delivery.ProductColorResponse {
			if len(p.Colors) == 0 {
				return nil
			}
			colors := make([]delivery.ProductColorResponse, len(p.Colors))
			for i, c := range p.Colors {
				colors[i] = delivery.ProductColorResponse{
					ProductItemID: c.ProductItemID,
					Color:         c.Color,
					Photo:         c.Photo,
				}
			}
			return colors
		}(),
		Promotions: func() []delivery.ProductPromotionResponse {
			if len(p.Promotions) == 0 {
				return nil
			}
			promotions := make([]delivery.ProductPromotionResponse, len(p.Promotions))
			for i, promo := range p.Promotions {
				promotions[i] = delivery.ProductPromotionResponse{
					ProductID:        promo.ProductID,
					Description:      promo.Description,
					DiscountPct:      promo.DiscountPct,
					Installments:     promo.Installments,
					InstallmentValue: promo.InstallmentValue,
					CardType:         promo.CardType,
					ExtraDiscountPct: promo.ExtraDiscountPct,
				}
			}
			return promotions
		}(),
		RelatedProducts: func() []delivery.RelatedProductsResponse {
			if len(p.Product.RelatedProducts) == 0 {
				return nil
			}
			related := make([]delivery.RelatedProductsResponse, len(p.Product.RelatedProducts))
			for i, rp := range p.Product.RelatedProducts {
				related[i] = delivery.RelatedProductsResponse{
					ProductItemID: rp.ProductItemID,
					Name:          rp.Name,
					Photo:         rp.Photo,
					Price:         rp.Price,
					Promotion:     rp.Promotion,
				}
			}
			return related
		}(),
		Seller: delivery.SellerResponse{
			SellerID:     p.SellerID,
			Name:         p.Seller.Name,
			OfficialSeal: p.Seller.OfficialSeal,
			Rating:       p.Seller.Rating,
		},
		PaymentMethods: func() []delivery.PaymentMethodResponse {
			if len(pm) == 0 {
				return nil
			}
			paymentMethods := make([]delivery.PaymentMethodResponse, len(pm))
			for i, method := range pm {
				paymentMethods[i] = delivery.PaymentMethodResponse{
					PaymentMethodID: method.PaymentMethodID,
					Icon:            method.Icon,
					Type:            method.Type,
				}
			}
			return paymentMethods
		}(),
	}
}
