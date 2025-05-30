package domain

type PaymentMethod struct {
	PaymentMethodID string
	Icon            string
	Type            string // "debit" or "credit"
}
