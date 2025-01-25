package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/pkg/db"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const collection = "orders"

func setDBindexes(collection *db.Collection) {
	collection.Indexes().
		CreateMany(context.Background(), []mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "id", Value: -1}},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys: bson.D{{Key: "customer.id", Value: -1}},
			},
		})
}

type Repository struct {
	db *db.Collection
}

func NewRepository(deps Deps) *Repository {
	return &Repository{db: deps.DB.Collection(collection)}
}

func (r *Repository) Persist(ctx context.Context, od Order) error {
	criteria := db.Criteria{
		Collection: collection,
		Filter:     bson.D{{Key: "id", Value: od.ID}},
	}

	return r.db.Persist(ctx, criteria, od)
}

func (r *Repository) Query(ctx context.Context, id uuid.UUID) db.Result[Order] {
	criteria := db.Criteria{
		Collection: collection,
		Filter:     bson.D{{Key: "id", Value: id}},
	}

	var res []Order
	if err := r.db.QueryMulti(ctx, criteria, &res); err != nil {
		return db.NewResult[Order](nil, fmt.Errorf("finding many documents: %w", err))
	}

	return db.NewResult(res, nil)
}
