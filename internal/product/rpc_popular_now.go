package product

import (
	"context"
	"fmt"

	"github.com/joseluis8906/pocone/pkg/log"
)

func (s *RpcService) PopularNow() ([]Product, error) {
	products, err := s.Repository.Query(context.Background(), Product{PopularNow: true}).All()
	if err != nil {
		err := fmt.Errorf("getting special offer: %w", err)
		log.Printf("%s %v", log.Error, err)
		return nil, err
	}

	return products, nil
}
