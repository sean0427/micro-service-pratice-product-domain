package product

import (
	"context"
	"testing"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type mockRepo struct{}

func (r *mockRepo) Get(ctx context.Context) ([]*model.Product, error) {
	return []*model.Product{{ID: "test"}}, nil
}

func (r *mockRepo) GetByID(ctx context.Context, id string) (*model.Product, error) {
	return &model.Product{
		ID:   "testfjeia",
		Name: id}, nil
}

func createMockRepo() *mockRepo {
	// TODO
	return &mockRepo{}
}

var testService *ProductService

func TestMain(m *testing.M) {
	testService = New(createMockRepo())
}

func TestProductService_Get(t *testing.T) {
	t.Run("Should success get product", func(t *testing.T) {
		list, err := testService.Get(context.TODO())

		if len(list) == 0 {
			t.Errorf("Get product list is empty")
		}

		if err != nil {
			t.Error(err)
		}
	})
}

func TestProductService_GetByID(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		const testName = "test"
		item, err := testService.GetByID(context.TODO(), testName)
		if err != nil {
			t.Error(err)
		}

		if item.Name != testName {
			t.Errorf("Get product by name is not equal")
		}

		if item.ID == "" {
			t.Errorf("Returned product id should not be empty")
		}
	})
}
