package customer

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *RpcService) Favorites(id uuid.UUID) ([]string, error) {
	c, err := s.Repository.Query(context.Background(), Customer{ID: id}).ExpectOne()
	if err != nil {
		return nil, fmt.Errorf("quering customer repository: %w", err)
	}

	return c.Favorites, nil
}
