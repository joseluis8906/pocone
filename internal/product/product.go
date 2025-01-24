package product

import (
	"github.com/joseluis8906/pocone/pkg/money"
)

type Product struct {
	Code         string
	Name         string
	Description  string
	Price        money.Money
	Discount     float64
	Image        string
	Categories   []string
	SpecialOffer bool
	PopularNow   bool
}

func PriceBeforeDiscount(p Product) money.Money {
	amount := float64(p.Price.Amount()) / 100
	x := amount / (1 - (p.Discount / 100))
	return money.NewFromFloat(x, p.Code)
}

var validCategories = map[string]struct{}{
	"meals":     {},
	"healthy":   {},
	"drinks":    {},
	"sides":     {},
	"snacks":    {},
	"fast food": {},
}
