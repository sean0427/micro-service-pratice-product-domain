package product

import (
	"testing"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type mockRepo struct{}

func (r *mockRepo) Get() ([]*model.Product, error) {
	return []*model.Product{{ID: "test"}}, nil
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
		list, err := testService.Get()

		if len(list) == 0 {
			t.Errorf("Get product list is empty")
		}

		if err != nil {
			t.Error(err)
		}
	})
}
