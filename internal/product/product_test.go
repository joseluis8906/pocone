package product_test

import (
	"testing"

	"github.com/joseluis8906/pocone/internal/product"
	"github.com/joseluis8906/pocone/pkg/money"
)

func TestCalcDiscount(t *testing.T) {
	testCases := map[string]struct {
		input product.Product
		want  money.Money
	}{
		"100 USD": {
			input: product.Product{
				Price:    money.USD(90, 0),
				Discount: 10.0,
			},
			want: money.USD(100, 0),
		},
		"50 USD": {
			input: product.Product{
				Price:    money.USD(37, 50),
				Discount: 25.0,
			},
			want: money.USD(50, 0),
		},
		"19500 COP": {
			input: product.Product{
				Price:    money.COP(13065, 0),
				Discount: 33.0,
			},
			want: money.COP(19500, 0),
		},
		"30000 COP": {
			input: product.Product{
				Price:    money.COP(30000, 0),
				Discount: 0,
			},
			want: money.COP(30000, 0),
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := product.OriginalPrice(tc.input)
			if tc.want.Amount() != got.Amount() {
				t.Errorf("want = %v; got = %v", tc.want.Amount(), got.Amount())
			}
		})

	}
}
