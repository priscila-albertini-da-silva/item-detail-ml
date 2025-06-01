package repository

import "github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"

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
