package product

import (
	"context"
	"errors"
	"fmt"

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
				Keys: bson.D{{Key: "specialoffer", Value: -1}},
			},
			{
				Keys: bson.D{{Key: "popularnow", Value: -1}},
			},
		})
}

type Repository struct {
	db *db.Collection
}

func (r *Repository) persist(ctx context.Context, p Product) error {
	filter := bson.D{{Key: "code", Value: p.Code}}
	opts := options.Replace().SetUpsert(true)

	_, err := db.DB.Collection(collection).ReplaceOne(ctx, filter, p, opts)

	return err
}

func (r *Repository) query(ctx context.Context, p Product) (res []Product, err error) {
	criteria := db.Criteria{
		Collection: collection,
	}

	var filter bson.M
	switch {
	case p.Code != "" && p.Name != "":
		filter = bson.M{"$or": []bson.M{
			{"code": bson.M{"$regex": p.Code, "$options": "i"}},
			{"Name": bson.M{"$regex": p.Name, "$options": "i"}},
		}}
	case p.Code != "":
		filter = bson.M{"code": bson.M{"$regex": p.Code, "$options": "i"}}
	case p.Name != "":
		filter = bson.M{"name": bson.M{"$regex": p.Name, "$options": "i"}}
	case len(p.Categories) > 0:
		filter = bson.M{"categories": bson.M{"$regex": p.Categories[0], "$options": "i"}}
	case p.SpecialOffer:
		filter = bson.M{"specialoffer": true}
	case p.PopularNow:
		filter = bson.M{"popularnow": true}
	}

	if len(filter) == 0 {
		return nil, errors.New("empty filter")
	}

	criteria.Filter = filter
	if err = db.QueryMulti(ctx, criteria, &res); err != nil {
		return nil, fmt.Errorf("quering mongo: %w", err)
	}

	return res, nil
}
