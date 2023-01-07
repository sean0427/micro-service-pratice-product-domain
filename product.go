package product

import (
	"context"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type repository interface {
	Get() ([]*model.Product, error)
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

func (s *ProductService) Get() ([]*model.Product, error) {
	return s.repo.Get()
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*model.Product, error) {
	return s.repo.GetByID(ctx, id)
}
