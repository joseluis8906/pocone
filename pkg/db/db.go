package db

import (
	"context"
	"errors"

	"github.com/joseluis8906/pocone/pkg/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/fx"
)

var DB *Database

type Deps struct {
	fx.In
	Config *viper.Viper
}

type (
	Database struct {
		*mongo.Database
		CollectionFn func(name string, opts ...options.Lister[options.CollectionOptions]) *Collection
	}

	Collection struct {
		*mongo.Collection
		ReplaceOneFn func(ctx context.Context, filter interface{}, replacement interface{}, opts ...options.Lister[options.ReplaceOptions]) (*UpdateResult, error)
		FindOneFn    func(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOneOptions]) *SingleResult
		FindFn       func(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOptions]) (*Cursor, error)
	}

	Cursor struct {
		*mongo.Cursor
		AllFn    func(ctx context.Context, results interface{}) error
		DecodeFn func(val interface{}) error
	}

	UpdateResult struct {
		*mongo.UpdateResult
	}

	SingleResult struct {
		*mongo.SingleResult
	}
)

func NewSingleResult(s *mongo.SingleResult) *SingleResult {
	return &SingleResult{s}
}

func NewUpdateResult(u *mongo.UpdateResult) *UpdateResult {
	return &UpdateResult{u}
}

func New(deps Deps) *Database {
	client, err := mongo.Connect(options.Client().ApplyURI(deps.Config.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("%s connecting mongo: %v", log.Error, err)
		return nil
	}

	mdb := client.Database("pocone")
	DB = &Database{
		Database: mdb,
	}
	return DB
}

func (c *Database) Collection(name string, opts ...options.Lister[options.CollectionOptions]) *Collection {
	if c.CollectionFn != nil {
		return c.CollectionFn(name, opts...)
	}

	return &Collection{Collection: c.Database.Collection(name, opts...)}
}

type Criteria struct {
	Collection string
	Filter     interface{}
}

func (c *Collection) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...options.Lister[options.ReplaceOptions]) (*UpdateResult, error) {
	if c.ReplaceOneFn != nil {
		return c.ReplaceOneFn(ctx, filter, replacement, opts...)
	}

	r, err := c.Collection.ReplaceOne(ctx, filter, replacement, opts...)

	return &UpdateResult{r}, err
}

func (c *Collection) FindOne(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOneOptions]) *SingleResult {
	if c.FindOneFn != nil {
		return c.FindOneFn(ctx, filter, opts...)
	}

	return &SingleResult{c.Collection.FindOne(ctx, filter, opts...)}
}

func (c *Collection) Find(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOptions]) (*Cursor, error) {
	if c.FindFn != nil {
		return c.FindFn(ctx, filter, opts...)
	}

	cur, err := c.Collection.Find(ctx, filter, opts...)
	return &Cursor{Cursor: cur}, err
}

func (c *Collection) Persist(ctx context.Context, criteria Criteria, v interface{}) error {
	opts := options.Replace().SetUpsert(true)
	_, err := DB.Collection(criteria.Collection).ReplaceOne(ctx, criteria.Filter, v, opts)

	return err
}

func (c *Collection) QuerySingle(ctx context.Context, criteria Criteria, v interface{}) error {
	ret := DB.Collection(criteria.Collection).FindOne(ctx, criteria.Filter)
	if err := ret.Err(); err != nil {
		return err
	}

	return ret.Decode(v)
}

func (c *Collection) QueryMulti(ctx context.Context, criteria Criteria, v interface{}) error {
	cur, err := DB.Collection(criteria.Collection).Find(ctx, criteria.Filter)
	if err != nil {
		return err
	}

	return cur.All(ctx, v)
}

type Result[T any] struct {
	data []T
	err  error
}

func NewResult[T any](data []T, err error) Result[T] {
	return Result[T]{data, err}
}

func (r Result[T]) ExpectOne() (*T, error) {
	if r.err != nil {
		return nil, r.err
	}

	if len(r.data) != 1 {
		return nil, errors.New("result length different to one")
	}

	v := r.data[0]
	return &v, nil
}

func (r Result[T]) ExpectMany() ([]T, error) {
	if len(r.data) == 0 {
		return nil, errors.New("empty result")
	}

	return r.data, nil
}

func (r Result[T]) All() ([]T, error) {
	return r.data, nil
}
