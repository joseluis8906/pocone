package customer

import (
	"context"
	"errors"
	"fmt"

	"github.com/joseluis8906/pocone/pkg/db"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const collection = "customers"

func setDBIndexes(collection *db.Collection) {
	collection.Indexes().
		CreateMany(context.Background(), []mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "id", Value: -1}},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys: bson.D{{Key: "name", Value: -1}},
			},
		})
}

type Repository struct {
	db *db.Collection
}

func NewRepository(deps Deps) *Repository {
	return &Repository{db: deps.DB.Collection(collection)}
}

func (r *Repository) Persist(ctx context.Context, c *Customer) error {
	filter := bson.D{{Key: "id", Value: c.ID}}
	opts := options.Replace().SetUpsert(true)

	if _, err := r.db.ReplaceOne(ctx, filter, c, opts); err != nil {
		return fmt.Errorf("replacing one document: %w", err)
	}

	return nil
}

func (r *Repository) Query(ctx context.Context, c Customer) db.Result[Customer] {
	criteria := db.Criteria{
		Collection: collection,
	}

	var filter bson.M
	switch {
	case c.ID.String() != "":
		filter = bson.M{"id": bson.M{"$regex": c.ID, "$options": "i"}}
	}

	if len(filter) == 0 {
		return db.NewResult[Customer](nil, errors.New("empty filter"))
	}

	criteria.Filter = filter
	var res []Customer
	if err := r.db.QueryMulti(ctx, criteria, &res); err != nil {
		return db.NewResult[Customer](nil, fmt.Errorf("finding many documents: %w", err))
	}

	return db.NewResult(res, nil)
}
