package product

import (
	"context"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type repository interface {
	Get(ctx context.Context) ([]*model.Product, error)
	GetByID(ctx context.Context, id string) (*model.Product, error)
}

type ProductService struct {
	repo repository
}

func New(repo repository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) Get(ctx context.Context) ([]*model.Product, error) {
	return s.repo.Get(ctx)
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*model.Product, error) {
	return s.repo.GetByID(ctx, id)
}
