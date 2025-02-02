package product

import (
	"context"
	"fmt"

	"github.com/joseluis8906/pocone/pkg/log"
)

func (s *RpcService) Search(ctx context.Context, req Product) ([]Product, error) {
	products, err := s.Repository.Query(context.Background(), req).All()
	if err != nil {
		err := fmt.Errorf("getting products: %w", err)
		log.Printf("%s %v", log.Error, err)
		return nil, err
	}

	return products, nil
}
