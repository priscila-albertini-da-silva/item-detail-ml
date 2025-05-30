package usecase

import "github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"

var MockPaymentMethod = domain.PaymentMethod{
	PaymentMethodID: "pm1",
	Icon:            "visa.png",
	Type:            "credit",
}

var MockPaymentMethods = []domain.PaymentMethod{
	{
		PaymentMethodID: "pm1",
		Icon:            "visa.png",
		Type:            "credit",
	},
	{
		PaymentMethodID: "pm2",
		Icon:            "mastercard.png",
		Type:            "debit",
	},
}
