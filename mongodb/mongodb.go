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
	client mongoClient
}

type mongoClient interface {
	Find(context.Context, interface{}, ...*options.FindOptions) (cur *mongo.Cursor, err error)
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) *mongo.SingleResult
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

var _ mongoClient = (*mongo.Collection)(nil)

const collectionName = "product_domain"

func all[T any](ctx context.Context, cursor *mongo.Cursor) (*T, error) {
	var results T
	if err := cursor.All(ctx, &results); err != nil {
		return nil, NewMongoError(err)
	}

	return &results, nil
}

// TODO: workaround fix for var function not support generic type
var productAll func(ctx context.Context, cursor *mongo.Cursor) (*[]*model.Product, error) = all[[]*model.Product]

func New(db *mongo.Database) *mongoRepository {
	return &mongoRepository{
		client: db.Collection(collectionName),
	}
}

func (m *mongoRepository) Get(ctx context.Context, params *model.GetProductsParams) ([]*model.Product, error) {
	b := bson.D{}
	if params.Name != nil {
		b = append(b, bson.E{Key: "name", Value: *params.Name})
	}
	if params.ManufacturerID != nil {
		b = append(b, bson.E{Key: "manufacturer_id", Value: *params.ManufacturerID})
	}
	cursor, err := m.client.Find(ctx, b)

	if err != nil {
		return nil, NewMongoError(err)
	}

	results, err := productAll(ctx, cursor)
	if err != nil {
		return nil, NewMongoError(err)
	}

	return *results, nil
}

func (m *mongoRepository) GetByID(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	err := m.client.FindOne(ctx, bson.E{Key: "_id", Value: id}).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return nil, nil // not found
	} else if err != nil {
		return nil, NewMongoError(err)
	}

	return &product, nil
}

func (m *mongoRepository) Create(ctx context.Context, params *model.Product) (string, error) {
	r, err := m.client.InsertOne(ctx, *params)
	if err != nil {
		return "", NewMongoError(err)
	}

	if v, ok := r.InsertedID.(primitive.ObjectID); ok {
		return v.Hex(), nil
	}

	return "", errors.New("invalid inserted id")
}

func (m *mongoRepository) Update(ctx context.Context, id string, params *model.Product) (*model.Product, error) {
	r, err := m.client.UpdateOne(ctx, bson.D{{Key: "_id", Value: id}}, *params)
	if err != nil {
		return nil, NewMongoError(err)
	}

	// TODO
	if r.ModifiedCount > 0 {
		return params, nil
	}
	return nil, nil
}
