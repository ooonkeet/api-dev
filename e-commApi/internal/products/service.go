package products

import (
	"context"

	repo "github.com/ooonkeet/api-dev/e-commApi/internal/adapters/postgresql/sqlc"
	// "github.com/ooonkeet/api-dev/e-commApi/internal/products"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product,error)
}

type svc struct{
	// repo
	repo repo.Querier
}

func NewService(repo repo.Querier) Service{
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product,error){
	// return nil
	return s.repo.ListProducts(ctx)

}
