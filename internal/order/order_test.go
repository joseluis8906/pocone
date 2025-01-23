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

	type Mongo struct {
		Data *db.UpdateResult
		Err  error
	}

	testCases := map[string]struct {
		mongo Mongo
		want  error
	}{
		"no errors": {
			mongo: Mongo{
				Data: nil,
				Err:  nil,
			},
			want: nil,
		},
		"db fails": {
			mongo: Mongo{
				Data: nil,
				Err:  testError,
			},
			want: testError,
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			mCol := mongotest.NewCollection()
			mDB := mongotest.New(mCol)
			db.DB = mDB

			mCol.ReplaceOneFn = func(context.Context, interface{}, interface{}, ...options.Lister[options.ReplaceOptions]) (*db.UpdateResult, error) {
				return tc.mongo.Data, tc.mongo.Err
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
	testDate := time.Date(2025, time.January, 21, 14, 12, 30, 0, time.UTC)
	testOrder := order.Order{
		ID:   testID,
		Date: testDate,
	}

	type (
		Input struct {
			Req *http.Request
		}
		Mongo struct {
			Data bson.M
			Err  error
		}
		Want struct {
			Data []byte
			Err  bool
		}
	)

	testCases := map[string]struct {
		input Input
		mongo Mongo
		want  Want
	}{
		"success": {
			input: Input{
				Req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "/orders/"+testID.String(), nil)
					r.SetPathValue("id", testID.String())
					return r
				}(),
			},
			mongo: Mongo{
				Data: bson.M{"id": testID, "date": testDate},
				Err:  nil,
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
				Req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "/orders/"+testID.String(), nil)
					r.SetPathValue("id", "abc-123")
					return r
				}(),
			},
			mongo: Mongo{
				Data: bson.M{},
				Err:  nil,
			},
			want: Want{
				Data: nil,
				Err:  true,
			},
		},
		"db fails": {
			input: Input{
				Req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "/orders/"+testID.String(), nil)
					r.SetPathValue("id", testID.String())
					return r
				}(),
			},
			mongo: Mongo{
				Data: bson.M{},
				Err:  testError,
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
				return mongotest.NewSingleResult(tc.mongo.Data, tc.mongo.Err)
			}

			var task order.GetTask
			task.ExtractID(tc.input.Req)
			task.SearchTheDB(ctx)
			task.Encode()

			got, err := task.Result()
			if diff := cmp.Diff(tc.want.Data, got); (err != nil && !tc.want.Err) || diff != "" {
				t.Errorf("order.GetTask() = %s, %v; want %s, %v\n(-want, +got)\n%s", got, err, tc.want.Data, testError, diff)
			}
		})
	}
}
