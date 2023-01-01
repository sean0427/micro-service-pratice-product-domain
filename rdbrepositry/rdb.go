package rdbreposity

import (
	"time"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type repository struct {
}

func New() *repository {
	return &repository{}
}

// TODO
func (r *repository) Get() ([]*model.Product, error) {
	return []*model.Product{
		{
			ID:             "mock1",
			Name:           "mock1",
			ManufacturerID: "test",
			CreatedBy:      "test",
			Created:        time.Now(),
			Updated:        time.Now(),
		},
		{
			ID:             "mock2",
			Name:           "mock2",
			ManufacturerID: "test",
			CreatedBy:      "test",
			Created:        time.Now(),
			Updated:        time.Now(),
		},
	}, nil
}
