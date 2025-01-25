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

func NewRepository(deps Deps) *Repository {
	return &Repository{db: deps.DB.Collection(collection)}
}

func (r *Repository) Persist(ctx context.Context, p *Product) error {
	filter := bson.D{{Key: "code", Value: p.Code}}
	opts := options.Replace().SetUpsert(true)

	if _, err := r.db.ReplaceOne(ctx, filter, p, opts); err != nil {
		return fmt.Errorf("calling replaceOne: %w", err)
	}

	return nil
}

func (r *Repository) Query(ctx context.Context, p Product) db.Result[Product] {
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
		return db.NewResult[Product](nil, errors.New("empty filter"))
	}

	criteria.Filter = filter
	var res []Product
	if err := r.db.QueryMulti(ctx, criteria, &res); err != nil {
		return db.NewResult[Product](nil, fmt.Errorf("finding many documents: %w", err))
	}

	return db.NewResult(res, nil)
}
