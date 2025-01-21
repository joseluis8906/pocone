package mongotest

import (
	"github.com/joseluis8906/pocone/pkg/db"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func New(col *db.Collection) *db.Database {
	return &db.Database{
		CollectionFn: func(string, ...options.Lister[options.CollectionOptions]) *db.Collection {
			return col
		},
	}
}

func NewCollection() *db.Collection {
	return &db.Collection{}
}

func NewSingleResult(v interface{}, err error) *db.SingleResult {
	return db.NewSingleResult(mongo.NewSingleResultFromDocument(v, err, bson.NewRegistry()))
}

func NewCursor() *db.Cursor {
	return &db.Cursor{}
}
