package factory

import (
	"reflect"
	"testing"

	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/factory"
)

func TestToItemProductResponse_AllFields(t *testing.T) {
	factory := factory.NewItemProductResponseFactory()

	product := MockProduct
	productItem := MockProductItem
	paymentMethods := MockPaymentMethods

	resp := factory.ToItemProductResponse(productItem, paymentMethods)

	if resp.ProductID != productItem.ProductID {
		t.Errorf("ProductID mismatch")
	}
	if resp.Price != productItem.Price {
		t.Errorf("Price mismatch")
	}
	if resp.PreviousPrice != productItem.PreviousPrice {
		t.Errorf("PreviousPrice mismatch")
	}
	if resp.Stock != productItem.Stock {
		t.Errorf("Stock mismatch")
	}
	if !reflect.DeepEqual(resp.Photos, productItem.Photos) {
		t.Errorf("Photos mismatch")
	}
	if resp.BestSeller != productItem.BestSeller {
		t.Errorf("BestSeller mismatch")
	}
	if resp.Description != product.Description {
		t.Errorf("Description mismatch")
	}
	if resp.Name != product.Name {
		t.Errorf("Name mismatch")
	}
	if resp.Brand != product.Brand {
		t.Errorf("Brand mismatch")
	}
	if !reflect.DeepEqual(resp.TechnicalFeatures, product.TechnicalFeatures) {
		t.Errorf("TechnicalFeatures mismatch")
	}
	if resp.Category.CategoryID != product.Category.CategoryID {
		t.Errorf("CategoryID mismatch")
	}
	if resp.Category.Description != product.Category.Description {
		t.Errorf("Category Description mismatch")
	}
	if resp.Category.ParentCategory == nil || resp.Category.ParentCategory.CategoryID != product.Category.ParentCategory.CategoryID {
		t.Errorf("ParentCategory mismatch")
	}
	if len(resp.Reviews) != 1 || resp.Reviews[0].ReviewID != productItem.Reviews[0].ReviewID {
		t.Errorf("Reviews mismatch")
	}
	if len(resp.Colors) != 1 || resp.Colors[0].Color != productItem.Colors[0].Color {
		t.Errorf("Colors mismatch")
	}
	if len(resp.Promotions) != 1 || resp.Promotions[0].Description != productItem.Promotions[0].Description {
		t.Errorf("Promotions mismatch")
	}
	if len(resp.RelatedProducts) != 1 || resp.RelatedProducts[0].ProductItemID != product.RelatedProducts[0].ProductItemID {
		t.Errorf("RelatedProducts mismatch")
	}
	if resp.Seller.SellerID != productItem.SellerID {
		t.Errorf("SellerID mismatch")
	}
	if resp.Seller.Name != productItem.Seller.Name {
		t.Errorf("Seller Name mismatch")
	}
	if resp.Seller.OfficialSeal != productItem.Seller.OfficialSeal {
		t.Errorf("Seller OfficialSeal mismatch")
	}
	if resp.Seller.Rating != productItem.Seller.Rating {
		t.Errorf("Seller Rating mismatch")
	}
	if len(resp.PaymentMethods) != 2 || resp.PaymentMethods[0].PaymentMethodID != paymentMethods[0].PaymentMethodID {
		t.Errorf("PaymentMethods mismatch")
	}
}

func TestToItemProductResponse_EmptySlicesAndNilParent(t *testing.T) {
	factory := factory.NewItemProductResponseFactory()

	productItem := MockProductItemEmptySlices

	resp := factory.ToItemProductResponse(productItem, nil)

	if len(resp.Photos) != 0 {
		t.Errorf("Expected nil Photos")
	}
	if resp.Reviews != nil {
		t.Errorf("Expected nil Reviews")
	}
	if resp.Colors != nil {
		t.Errorf("Expected nil Colors")
	}
	if resp.Promotions != nil {
		t.Errorf("Expected nil Promotions")
	}
	if resp.RelatedProducts != nil {
		t.Errorf("Expected nil RelatedProducts")
	}
	if resp.Category.ParentCategory != nil {
		t.Errorf("Expected nil ParentCategory")
	}
	if resp.PaymentMethods != nil {
		t.Errorf("Expected nil PaymentMethods")
	}
}
