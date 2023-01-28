package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type mongoRepository struct {
	client MongoClient
}

type MongoClient interface {
	Find(context.Context, interface{}, ...*options.FindOptions) (cur *mongo.Cursor, err error)
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) *mongo.SingleResult
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

var _ MongoClient = (*mongo.Collection)(nil)

const collectionName = "product_domain"

func New(db *mongo.Database) *mongoRepository {
	return &mongoRepository{
		client: db.Collection(collectionName),
	}
}

func (m *mongoRepository) Get(ctx context.Context, params *model.GetProductsParams) ([]*model.Product, error) {
	cursor, err := m.client.Find(ctx, bson.D{{Key: "name", Value: params.Name},
		{Key: "Manufacturer", Value: params.ManufacturerID}})
	if err != nil {
		return nil, NewMongoError(err)
	}

	var results []*model.Product
	if err = cursor.All(ctx, &results); err != nil {
		return nil, NewMongoError(err)
	}

	return results, nil
}

func (m *mongoRepository) GetByID(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	err := m.client.FindOne(ctx, primitive.E{Key: "ID", Value: id}).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return nil, NewMongoError(err)
	}

	return &product, nil
}

func (m *mongoRepository) Create(ctx context.Context, params *model.Product) (string, error) {
	r, err := m.client.InsertOne(ctx, params)
	if err != nil {
		return "", NewMongoError(err)
	}

	if v, ok := r.InsertedID.(primitive.ObjectID); ok {
		return v.String(), nil
	}

	return "", errors.New("invalid inserted id")
}

func (m *mongoRepository) Update(ctx context.Context, id string, params *model.Product) (*model.Product, error) {
	r, err := m.client.UpdateOne(ctx, bson.D{{Key: "ID", Value: id}}, bson.M{"$set": params})
	if err != nil {
		return nil, NewMongoError(err)
	}

	// TODO
	if r.ModifiedCount > 0 {
		return params, nil
	}
	return nil, nil
}
