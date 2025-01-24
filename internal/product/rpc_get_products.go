package product

import (
	"context"
	"fmt"

	"github.com/joseluis8906/pocone/pkg/log"
)

type (
	GetProductsReply struct {
		Products []Product
	}
)

func (s *RpcService) Get(req *Product) ([]Product, error) {
	products, err := s.Repository.query(context.Background(), *req)
	if err != nil {
		log.Printf("%s getting products: %v", log.Error, err)
		return nil, fmt.Errorf("getting products: %w", err)
	}

	return products, nil
}
