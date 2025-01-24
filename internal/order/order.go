package order

import (
	"time"

	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/internal/customer"
	"github.com/joseluis8906/pocone/internal/product"
	"github.com/joseluis8906/pocone/pkg/money"
)

type (
	Order struct {
		ID       uuid.UUID
		Date     time.Time
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
