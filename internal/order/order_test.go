package order_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/internal/order"
	"github.com/joseluis8906/pocone/pkg/db"
	"github.com/joseluis8906/pocone/pkg/mongotest"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestCreateTask(t *testing.T) {
	ctx := context.Background()
	testError := errors.New("test error")

	testCases := map[string]struct {
		input error
		want  error
	}{
		"no errors": {},
		"db fails": {
			input: testError,
			want:  testError,
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			mCol := mongotest.NewCollection()
			mDB := mongotest.New(mCol)
			db.DB = mDB

			mCol.ReplaceOneFn = func(context.Context, interface{}, interface{}, ...options.Lister[options.ReplaceOptions]) (*db.UpdateResult, error) {
				return nil, tc.input
			}

			var task order.CreateTask
			task.GenerateID()
			task.PersistOrder(ctx)
			task.MarshalOrder()

			_, got := task.Result()
			if !errors.Is(got, tc.want) {
				t.Errorf("order.CreateTask() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetTask(t *testing.T) {
	ctx := context.Background()
	testError := errors.New("test error")

	testID := uuid.New()
	testOrder := order.Order{
		ID:   testID,
		Date: time.Date(2025, time.January, 21, 14, 12, 30, 0, time.UTC),
	}

	type (
		Input struct {
			Order order.Order
			ReqID string
			Err   error
		}
		Want struct {
			Data []byte
			Err  bool
		}
	)

	testCases := map[string]struct {
		input Input
		want  Want
	}{
		"success": {
			input: Input{
				Order: testOrder,
				ReqID: testID.String(),
				Err:   nil,
			},
			want: Want{
				Data: func() []byte {
					d, _ := json.Marshal(testOrder)
					return d
				}(),
				Err: false,
			},
		},
		"wrong req param id": {
			input: Input{
				Order: testOrder,
				ReqID: "wrong id format",
			},
			want: Want{
				Data: nil,
				Err:  true,
			},
		},
		"db fails": {
			input: Input{
				Order: testOrder,
				ReqID: testID.String(),
				Err:   testError,
			},
			want: Want{
				Data: nil,
				Err:  true,
			},
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			mCol := mongotest.NewCollection()
			db.DB = mongotest.New(mCol)

			mCol.FindOneFn = func(context.Context, interface{}, ...options.Lister[options.FindOneOptions]) *db.SingleResult {
				record := bson.M{
					"id":   tc.input.Order.ID,
					"date": tc.input.Order.Date,
				}

				return mongotest.NewSingleResult(record, tc.input.Err)
			}

			req := httptest.NewRequest(http.MethodGet, "/orders/"+tc.input.Order.ID.String(), nil)
			req.SetPathValue("id", tc.input.ReqID)

			var task order.GetTask
			task.ExtractID(req)
			task.SearchTheDB(ctx)
			task.Encode()

			got, err := task.Result()

			var want []byte
			if err == nil {
				want, _ = json.Marshal(tc.input.Order)
			}

			if diff := cmp.Diff(want, got); (err != nil && !tc.want.Err) || diff != "" {
				t.Errorf("order.GetTask() = %s, %v; want %s, %v\n(-want, +got)\n%s", got, err, tc.want.Data, testError, diff)
			}
		})
	}
}
