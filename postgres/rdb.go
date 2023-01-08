package rdbreposity

import (
	"context"

	"gorm.io/gorm"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, params model.GetProductsParams) ([]*model.Product, error) {
	// TODO
	var products []*model.Product
	tx := r.db.WithContext(ctx)

	if params.Name != nil {
		tx = tx.Where("name = ?", *params.Name)
	}

	result := tx.Find(&products)

	return products, result.Error
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	tx := r.db.WithContext(ctx)

	result := tx.Where("id =?", id).Find(&product)
	return &product, result.Error
}
