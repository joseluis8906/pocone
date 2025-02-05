package product_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/joseluis8906/pocone/internal/product"
	"github.com/joseluis8906/pocone/pkg/log"
	"github.com/joseluis8906/pocone/pkg/money"
	"github.com/joseluis8906/pocone/pkg/mongo"
	"github.com/joseluis8906/pocone/pkg/mongo/mongotest"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestRpcService_Add(t *testing.T) {
	log.Noop()

	testCases := map[string]struct {
		input    product.Product
		want     error
		mongoErr error
	}{
		"happy path": {
			input: product.Product{
				Code:        "ABC-123",
				Name:        "Cheese Burger",
				Description: "Classic cheese burger.",
				Price:       money.USD(5, 50),
				Discount:    0,
				Image:       "http://cdn.example.com/myimage.jpg",
				Categories:  []string{"fast food"},
			},
			want:     nil,
			mongoErr: nil,
		},
		"validation fails": {
			input: product.Product{
				Name:        "Cheese Burger",
				Description: "Classic cheese burger.",
				Price:       money.USD(5, 50),
				Discount:    0,
				Image:       "http://cdn.example.com/myimage.jpg",
				Categories:  []string{"fast food"},
			},
			want:     errors.New("validating product: empty code"),
			mongoErr: nil,
		},
		"mongo fails": {
			input: product.Product{
				Code:        "ABC-123",
				Name:        "Cheese Burger",
				Description: "Classic cheese burger.",
				Price:       money.USD(5, 50),
				Discount:    0,
				Image:       "http://cdn.example.com/myimage.jpg",
				Categories:  []string{"fast food"},
			},
			want:     errors.New("persisting product: calling replaceOne: mongo failed"),
			mongoErr: errors.New("mongo failed"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			replaceOneFn := mongotest.ReplaceOneFn(func(context.Context, interface{}, interface{}, ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error) {
				return mongotest.NewUpdateResult(nil), tc.mongoErr
			})

			colFn := mongotest.CollFn(func(string, ...options.Lister[options.CollectionOptions]) *mongo.Collection {
				return mongotest.NewCollection(replaceOneFn, nil, nil)
			})

			svc := product.RpcService{
				Repository: product.NewRepository(product.Deps{DB: mongotest.New(colFn)}),
			}

			got := svc.Add(context.Background(), &tc.input)
			if diff := cmp.Diff(fmt.Sprintf("%v", tc.want), fmt.Sprintf("%v", got)); diff != "" {
				t.Errorf("product.RpcService.Add(ctx, %v) = %v; want = %v, \n(-want, +got)\n%s", tc.input, got, tc.want, diff)
			}
		})
	}
}
