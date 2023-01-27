package rdbreposity

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type repository struct {
	client MongoClient
}

type MongoClient interface {
	Find(context.Context, interface{}, ...*options.FindOptions) (cur *mongo.Cursor, err error)
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) *mongo.SingleResult
}

const collectionName = "product_domain"

func New(db *mongo.Database) *repository {
	return &repository{
		client: db.Collection(collectionName),
	}
}

func (r *repository) Get(ctx context.Context, params *model.GetProductsParams) ([]*model.Product, error) {
	cursor, err := r.client.Find(ctx, bson.D{{Key: "name", Value: params.Name},
		{Key: "Manufacturer", Value: params.ManufacturerID}})
	if err != nil {
		return nil, err
	}

	var results []*model.Product
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	err := r.client.FindOne(ctx, primitive.E{Key: "ID", Value: id}).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	return &product, nil
}

var _ MongoClient = (*mongo.Collection)(nil)
