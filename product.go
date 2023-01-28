package product

import (
	"context"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type repository interface {
	Get(ctx context.Context, params *model.GetProductsParams) ([]*model.Product, error)
	GetByID(ctx context.Context, id string) (*model.Product, error)
	Create(ctx context.Context, product *model.Product) (string, error)
	Update(ctx context.Context, id string, params *model.Product) (*model.Product, error)
}

type ProductService struct {
	repo repository
}

func New(repo repository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) Get(ctx context.Context, params *model.GetProductsParams) ([]*model.Product, error) {
	return s.repo.Get(ctx, params)
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*model.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) Create(ctx context.Context, product *model.Product) (string, error) {
	return s.repo.Create(ctx, product)
}

func (s *ProductService) Update(ctx context.Context, id string, params *model.Product) (*model.Product, error) {
	return s.repo.Update(ctx, id, params)
}
