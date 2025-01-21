package order

import (
	"context"

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

func persistOnDB(ctx context.Context, od Order) error {
	criteria := db.Criteria{
		Collection: collection,
		Filter:     bson.D{{Key: "id", Value: od.ID}},
	}

	return db.Persist(ctx, criteria, od)
}

func queryByID(ctx context.Context, id uuid.UUID) (od Order, err error) {
	criteria := db.Criteria{
		Collection: collection,
		Filter:     bson.D{{Key: "id", Value: id}},
	}

	return od, db.QuerySingle(ctx, criteria, &od)
}
