package mongodb

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sean0427/micro-service-pratice-product-domain/mock"
	"github.com/sean0427/micro-service-pratice-product-domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FuzzMondoClient_Get(f *testing.F) {
	originalProductAll := productAll

	// Testing it not really care about the cursor
	mockCursor := &mongo.Cursor{}
	returnedResult := &[]*model.Product{
		{
			ID:             primitive.NewObjectID(),
			Name:           "test",
			ManufacturerID: primitive.NewObjectID(),
		},
	}

	for i := 0; i < 100; i++ {
		*returnedResult = append(*returnedResult, &model.Product{ID: primitive.NewObjectID(),
			Name:           uuid.NewString(),
			ManufacturerID: primitive.NewObjectID(),
			CreatedBy:      uuid.NewString(),
			Created:        time.Now(),
			Updated:        time.UnixMicro(int64(i)),
		})
	}

	f.Add("test", "", "", "")
	f.Add("test", "tes1t", "", "")
	f.Add("test", "tes2t", "error2", "")
	f.Add("test", "tes3t", "", "error")
	f.Add("test", "tes4t", "rror", "error")

	f.Fuzz(func(t *testing.T, name, manufacturer, errMsg, getCursorErrMsg string) {
		mc := mock.NewMockmongoClient(gomock.NewController(t))

		mc.EXPECT().
			Find(gomock.Any(), gomock.Eq(bson.D{{Key: "name", Value: name}, {Key: "manufacturer_id", Value: manufacturer}})).
			DoAndReturn(func(context.Context, interface{}, ...*options.FindOptions) (*mongo.Cursor, error) {
				if errMsg != "" {
					return nil, errors.New(errMsg)
				}
				return mockCursor, nil
			}).
			Times(1)

		productAll = func(ctx context.Context, cursor *mongo.Cursor) (*[]*model.Product, error) {
			if getCursorErrMsg != "" {
				return nil, errors.New(getCursorErrMsg)
			}

			if cursor != mockCursor {
				t.Fatalf("not expected cursor")
			}

			return returnedResult, nil
		}

		params := model.GetProductsParams{
			Name:           &name,
			ManufacturerID: &manufacturer,
		}

		client := mongoRepository{client: mc}
		ret, err := client.Get(context.Background(), &params)
		if errMsg != "" || getCursorErrMsg != "" {
			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if !errors.Is(&MongoError{}, err) {
				t.Errorf("error should be MongoError, but got %v", err)
			}

			unWrap := errors.Unwrap(err).Error()

			// errMsg first
			if errMsg != "" {
				if errMsg != unWrap {
					t.Errorf("expected error to be %s, got %s", errMsg, unWrap)
				}
				return
			}

			if getCursorErrMsg != "" && getCursorErrMsg != unWrap {
				t.Errorf("expected error to be %s, got %s", getCursorErrMsg, unWrap)
				return
			}

			return
		}

		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		if reflect.DeepEqual(returnedResult, ret) {
			t.Errorf("error, expected result to be %v, got %v", *returnedResult, ret)
		}
	})

	productAll = originalProductAll
}

func FuzzMongoClient_GetByID(f *testing.F) {
	f.Add("test", "")

	for i := 0; i < 100; i++ {
		f.Add(primitive.NewObjectID().String(), "")
	}

	f.Add("testErr", "err")

	returnedResult := model.Product{}
	f.Fuzz(func(t *testing.T, ID, errMsg string) {
		mc := mock.NewMockmongoClient(gomock.NewController(t))
		mc.EXPECT().
			FindOne(gomock.Any(), gomock.Eq(bson.E{Key: "_id", Value: ID})).
			DoAndReturn(func(context.Context, interface{}, ...*options.FindOneOptions) *mongo.SingleResult {
				if errMsg != "" {
					return mongo.NewSingleResultFromDocument(&model.Product{}, errors.New(errMsg), nil)
				}

				return mongo.NewSingleResultFromDocument(returnedResult, nil, nil)

			}).
			Times(1)

		client := mongoRepository{client: mc}
		ret, err := client.GetByID(context.Background(), ID)
		if errMsg != "" {
			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if !errors.Is(&MongoError{}, err) {
				t.Errorf("error should be MongoError, but got %v", err)
				return
			}
			return
		}

		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
		if *ret != returnedResult {
			t.Errorf("error, expected result to be %v, got %v", *ret, *ret)
		}
	})
}

func FuzzMongoClient_Create(f *testing.F) {
	for i := 0; i < 100; i++ {
		f.Add(primitive.NewObjectID().Hex(), "")
	}

	f.Add("test", "testErr")
	f.Add("test", "testErr2")

	f.Fuzz(func(t *testing.T, returnId, errMsg string) {
		testItem := model.Product{
			ID:   primitive.NewObjectID(),
			Name: uuid.NewString(),
		}
		gomock.NewController(t)

		mc := mock.NewMockmongoClient(gomock.NewController(t))
		mc.EXPECT().
			InsertOne(gomock.Any(), testItem).
			DoAndReturn(func(_ context.Context, params interface{}, _ ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
				if errMsg != "" {
					return nil, errors.New(errMsg)
				}
				if _, ok := params.(model.Product); ok {
					_id, _ := primitive.ObjectIDFromHex(returnId)
					return &mongo.InsertOneResult{InsertedID: _id}, nil
				}

				t.Fatalf("invalid type: %T", params)
				return nil, nil
			}).
			Times(1)

		client := mongoRepository{client: mc}
		id, err := client.Create(context.Background(), &testItem)
		if errMsg != "" {
			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if !errors.Is(&MongoError{}, err) {
				t.Errorf("error should be MongoError, but got %v", err)
				return
			}
			return
		}

		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
		if id != returnId {
			t.Errorf("expected id to be %s, got %s", returnId, id)
		}
	})
}

func FuzzMondoClient_Update(f *testing.F) {
	returnedResult := &mongo.UpdateResult{ModifiedCount: int64(1)}

	for i := 0; i < 100; i++ {
		f.Add(uuid.NewString(), uuid.NewString(), primitive.NewObjectID().String(), "")
	}
	f.Add("tewst", "tseta", "feafge", "")
	f.Add("tewst1", "tseta2", "feafge", "err")
	f.Add("tewst2", "tseta3", "feafge", "err2")

	f.Fuzz(func(t *testing.T, name, manufacturer, ID, errMsg string) {
		testItem := model.Product{Name: name, ManufacturerID: primitive.NewObjectID()}

		mc := mock.NewMockmongoClient(gomock.NewController(t))
		mc.EXPECT().
			UpdateOne(gomock.Any(), bson.D{{Key: "_id", Value: ID}}, testItem).
			DoAndReturn(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
				if errMsg != "" {
					return nil, errors.New(errMsg)
				}

				return returnedResult, nil
			}).
			Times(1)

		r := mongoRepository{client: mc}

		ret, err := r.Update(context.Background(), ID, &testItem)

		if errMsg != "" {
			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if !errors.Is(&MongoError{}, err) {
				t.Errorf("error should be MongoError, but got %v", err)
				return
			}
			unWrap := errors.Unwrap(err).Error()
			if errMsg != unWrap {
				t.Errorf("expected error to be %s, got %s", errMsg, unWrap)
				return
			}
			return
		}

		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		if ret != &testItem {
			t.Errorf("expected %v, got %v", &testItem, ret)
		}
	})
}
