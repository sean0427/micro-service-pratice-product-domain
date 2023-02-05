package product

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sean0427/micro-service-pratice-product-domain/mock"
	"github.com/sean0427/micro-service-pratice-product-domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FuzzProductService_Get(f *testing.F) {
	f.Add("test", "testm", "", 2)
	f.Add("test2", "testm2", "", 2)
	f.Add("test3", "testm5", "", 52)
	f.Add("test3", "testm4", "fe", 2)
	f.Add("test2", "testm5", "", 22)
	f.Add("test1", "testm2", "e", 2)
	f.Add("test1", "testm2", "e", 2)

	f.Fuzz(func(t *testing.T, name, manu, errMsg string, returnedProducts int) {
		c := gomock.NewController(t)
		repo := mock.NewMockrepository(c)

		request := &model.GetProductsParams{
			Name:           &name,
			ManufacturerID: &manu,
		}
		repo.EXPECT().
			Get(gomock.Any(), request).
			DoAndReturn(func(context.Context, *model.GetProductsParams) ([]*model.Product, error) {
				if errMsg != "" {
					return nil, errors.New(errMsg)
				}

				response := make([]*model.Product, returnedProducts)

				return response, nil
			}).
			Times(1)

		service := New(repo)

		list, err := service.Get(context.Background(), request)
		if errMsg != "" {
			if err == nil {
				t.Errorf("Error should not be nil")
			}

			if list != nil {
				t.Errorf("Get product should be nil when error")
			}
			return
		}

		if len(list) != returnedProducts {
			t.Errorf("Get product expect %d, %d", len(list), returnedProducts)
		}

	})
}

func FuzzProductService_GetByID(f *testing.F) {
	f.Add(primitive.NewObjectID().Hex(), "name2", primitive.NewObjectID().Hex(), "")
	f.Add(primitive.NewObjectID().Hex(), "name3", primitive.NewObjectID().Hex(), "")
	f.Add(primitive.NewObjectID().Hex(), "name52", primitive.NewObjectID().Hex(), "14")
	f.Add(primitive.NewObjectID().Hex(), "name42", primitive.NewObjectID().Hex(), "")
	f.Add(primitive.NewObjectID().Hex(), "name2", primitive.NewObjectID().Hex(), "313")

	f.Fuzz(func(t *testing.T, id, name, manu, errMsg string) {
		c := gomock.NewController(t)
		repo := mock.NewMockrepository(c)
		service := New(repo)

		repo.EXPECT().
			GetByID(gomock.Any(), id).
			DoAndReturn(func(ctx context.Context, id string) (*model.Product, error) {
				if errMsg != "" {
					return nil, errors.New(errMsg)
				}
				hex, _ := primitive.ObjectIDFromHex(id)
				man, _ := primitive.ObjectIDFromHex(manu)
				return &model.Product{
					ID:             hex,
					Name:           name,
					ManufacturerID: man,
				}, nil
			}).
			Times(1)

		item, err := service.GetByID(context.Background(), id)
		if errMsg != "" {
			if err == nil {
				t.Errorf("Error should not be nil")
			}

			if item != nil {
				t.Errorf("Get product should be nil when error")
			}
			return
		}

		if item.ID.Hex() != id {
			t.Errorf("Get product expect %s, %s", id, item.ID)
		}
		if item.Name != name {
			t.Errorf("Get product expect %s, %s", name, item.Name)
		}
		if item.ManufacturerID.Hex() != manu {
			t.Errorf("Get product expect %s, %s", manu, item.ManufacturerID)
		}
	})
}

func FuzzProductService_Update(f *testing.F) {
	f.Add("test", "name2", "naum", "")
	f.Add("test4", "name3", "naum", "")
	f.Add("test43", "name52", "naum", "14")
	f.Add("test4342", "name42", "naum", "")
	f.Add("test41", "name2", "naum", "313")

	f.Fuzz(func(t *testing.T, id, name, manu, errMsg string) {
		c := gomock.NewController(t)
		repo := mock.NewMockrepository(c)
		service := New(repo)

		params := &model.Product{
			Name:           name,
			ManufacturerID: primitive.NewObjectID(),
		}
		repo.EXPECT().
			Update(gomock.Any(), id, params).
			DoAndReturn(func(ctx context.Context, id string, params *model.Product) (*model.Product, error) {
				if errMsg != "" {
					return nil, errors.New(errMsg)
				}

				return params, nil
			}).
			Times(1)

		item, err := service.Update(context.Background(), id, params)
		if errMsg != "" {
			if err == nil {
				t.Errorf("Error should not be nil")
			}

			if item != nil {
				t.Errorf("Get product should be nil when error")
			}
			return
		}

		if item != params {
			t.Errorf("Get product expect %s, %s", id, item.ID)
		}
	})
}

func FuzzProductService_Create(f *testing.F) {
	f.Add("test", "name2", "naum", "")
	f.Add("test4", "name3", "naum", "")
	f.Add("test43", "name52", "naum", "14")
	f.Add("test4342", "name42", "naum", "")
	f.Add("test41", "name2", "naum", "313")

	f.Fuzz(func(t *testing.T, id, name, manu, errMsg string) {
		c := gomock.NewController(t)
		repo := mock.NewMockrepository(c)
		service := New(repo)

		params := &model.Product{
			Name:           name,
			ManufacturerID: primitive.NewObjectID(),
		}
		repo.EXPECT().
			Create(gomock.Any(), params).
			DoAndReturn(func(ctx context.Context, params *model.Product) (string, error) {
				if errMsg != "" {
					return "", errors.New(errMsg)
				}

				return id, nil
			}).
			Times(1)

		rID, err := service.Create(context.Background(), params)
		if errMsg != "" {
			if err == nil {
				t.Errorf("Error should not be nil")
			}

			if rID != "" {
				t.Errorf("Get product should be nil when error")
			}
			return
		}

		if rID != id {
			t.Errorf("Get product expect %s, %s", rID, id)
		}
	})
}
