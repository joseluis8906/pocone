package customer

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/internal/product"
)

func (s *RpcService) AddFavorite(id uuid.UUID, code string) error {
	ctx := context.Background()

	c, err := s.Repository.Query(ctx, Customer{ID: id}).ExpectOne()
	if err != nil {
		return fmt.Errorf("quering customer repository: %w", err)
	}

	p, err := s.ProductRepo.Query(ctx, product.Product{Code: code}).ExpectOne()
	if err != nil {
		return fmt.Errorf("quering product repository: %w", err)
	}

	c.Favorites = append(c.Favorites, p.Code)
	if err := s.Repository.Persist(ctx, c); err != nil {
		return fmt.Errorf("persisting customer: %w", err)
	}

	return nil
}
