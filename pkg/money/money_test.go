package money_test

import (
	"testing"

	"github.com/joseluis8906/pocone/pkg/money"
)

func TestAdd(t *testing.T) {
	type Input struct {
		A money.Money
		B money.Money
	}
	testCases := map[string]struct {
		input Input
		want  money.Money
	}{
		"10 USD": {
			input: Input{
				A: money.USD(7, 0),
				B: money.USD(3, 0),
			},
			want: money.USD(10, 0),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := money.Add(tc.input.A, tc.input.B)
			if tc.want.Display() != got.Display() {
				t.Errorf("money.Add(%v, %v) = %v; want = %v", tc.input.A.Display(), tc.input.B.Display(), got.Display(), tc.want.Display())
			}
		})
	}
}
