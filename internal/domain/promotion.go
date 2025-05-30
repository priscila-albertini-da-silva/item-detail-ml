package domain

type Promotion struct {
	ProductID        string
	Description      string
	DiscountPct      float64
	Installments     int
	InstallmentValue float64
	CardType         string
	ExtraDiscountPct float64
}
