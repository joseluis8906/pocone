package mongotest

import (
	"context"

	"github.com/joseluis8906/pocone/pkg/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
	stdmongo "go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type CollFn func(string, ...options.Lister[options.CollectionOptions]) *mongo.Collection

type (
	ReplaceOneFn func(ctx context.Context, filter interface{}, replacement interface{}, opts ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error)
	FindOneFn    func(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult
	FindFn       func(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error)
)

func New(collection CollFn) *mongo.Database {
	return &mongo.Database{
		CollectionFn: collection,
	}
}

func NewCollection(replaceOneFn ReplaceOneFn, findOneFn FindOneFn, findFn FindFn) *mongo.Collection {
	return &mongo.Collection{ReplaceOneFn: replaceOneFn, FindFn: findFn, FindOneFn: findOneFn}
}

func NewSingleResult(v interface{}, err error) *mongo.SingleResult {
	return mongo.NewSingleResult(stdmongo.NewSingleResultFromDocument(v, err, bson.NewRegistry()))
}

func NewUpdateResult(upsertedID interface{}) *mongo.UpdateResult {
	return mongo.NewUpdateResult(&stdmongo.UpdateResult{UpsertedID: upsertedID})
}

type (
	AllFn    = func(ctx context.Context, results interface{}) error
	DecodeFn = func(val interface{}) error
)

func NewCursor(allFn AllFn, decodeFn DecodeFn) *mongo.Cursor {
	return &mongo.Cursor{AllFn: allFn, DecodeFn: decodeFn}
}
