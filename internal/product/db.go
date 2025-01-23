package product

import (
	"context"

	"github.com/joseluis8906/pocone/pkg/db"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const collection = "products"

func setDBIndexes(collection *db.Collection) {
	collection.Indexes().
		CreateMany(context.Background(), []mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "code", Value: -1}},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys: bson.D{{Key: "name", Value: -1}},
			},
			{
				Keys: bson.D{{Key: "categories", Value: -1}},
			},
			{
				Keys: bson.D{{Key: "is_special_offer", Value: -1}},
			},
		})
}

func persistOnDB(ctx context.Context, p Product) error {
	filter := bson.D{{Key: "code", Value: p.Code}}
	opts := options.Replace().SetUpsert(true)

	_, err := db.DB.Collection(collection).ReplaceOne(ctx, filter, p, opts)

	return err
}

func queryDB(ctx context.Context, p Product) (res []Product, err error) {
	var filter bson.M
	criteria := db.Criteria{
		Collection: collection,
		Filter:     filter,
	}

	switch {
	case p.Code != "" && p.Name != "":
		filter = bson.M{"$or": []bson.M{
			{"code": bson.M{"$regex": p.Code}},
			{"Name": bson.M{"$regex": p.Name}},
		}}
	case p.Code != "":
		filter = bson.M{"code": bson.M{"$regex": p.Code}}
	case p.Name != "":
		filter = bson.M{"name": bson.M{"$regex": p.Name}}
	}

	return res, db.QueryMulti(ctx, criteria, &res)
}
