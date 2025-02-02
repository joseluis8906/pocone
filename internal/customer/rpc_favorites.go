package customer

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *RpcService) Favorites(ctx context.Context, id uuid.UUID) ([]string, error) {
	c, err := s.CustomerRepo.Query(ctx, Customer{ID: id}).ExpectOne()
	if err != nil {
		return nil, fmt.Errorf("quering customer repository: %w", err)
	}

	return c.Favorites, nil
}
