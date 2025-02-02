package order

import (
	"time"

	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/internal/address"
	"github.com/joseluis8906/pocone/internal/customer"
	"github.com/joseluis8906/pocone/internal/product"
	"github.com/joseluis8906/pocone/pkg/money"
)

type (
	Order struct {
		ID       uuid.UUID
		Date     time.Time
		Address  address.Address
		Customer customer.Customer
		Items    []Item
		Subtotal money.Money
		Fees     money.Money
		Discount float64
		Total    money.Money
	}

	Item struct {
		Product  product.Product
		Quantity int
		Price    money.Money
		Total    money.Money
	}
)

func CalculateTotal(ord Order) money.Money {
	var total money.Money
	for i, item := range ord.Items {
		if i == 0 {
			total = item.Price
			continue
		}

		total = money.Add(total, item.Price)
	}

	return total
}
