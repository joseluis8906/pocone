package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/pkg/log"
	"github.com/joseluis8906/pocone/pkg/slices"
)

func (s *RpcService) RemoveItem(ctx context.Context, orderID uuid.UUID, position int) (*Order, error) {
	ord, err := s.OrderRepo.Query(ctx, orderID).ExpectOne()
	if err != nil {
		err = fmt.Errorf("getting order: %w", err)
		log.Printf("%s %s", log.Error, err)

		return nil, err
	}

	index := position - 1
	if len(ord.Items) <= index {
		err := fmt.Errorf("error calculating index, length %d index %d: index out of range", len(ord.Items), index)
		log.Printf("%s %s", log.Error, err)
		return nil, err
	}

	ord.Items = slices.Splice(ord.Items, index)
	ord.Total = CalculateTotal(*ord)

	if err := s.OrderRepo.Persist(ctx, ord); err != nil {
		err = fmt.Errorf("persisting order: %w", err)
		log.Printf("%s %s", log.Error, err)

		return nil, err
	}

	return ord, nil
}
