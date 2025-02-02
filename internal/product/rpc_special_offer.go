package product

import (
	"context"
	"fmt"

	"github.com/joseluis8906/pocone/pkg/log"
	"github.com/joseluis8906/pocone/pkg/slices"
)

func (s *RpcService) SpecialOffer() ([]Product, error) {
	products, err := s.Repository.Query(context.Background(), Product{SpecialOffer: true}).All()
	if err != nil {
		err := fmt.Errorf("getting special offer: %w", err)
		log.Printf("%s %v", log.Error, err)
		return nil, err
	}

	return slices.Map(products, applyDiscount), nil
}
