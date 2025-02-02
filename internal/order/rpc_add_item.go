package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/internal/product"
	"github.com/joseluis8906/pocone/pkg/log"
)

func (s *RpcService) AddItem(ctx context.Context, orderID uuid.UUID, productID string, qty int) (*Order, error) {
	ord, err := s.OrderRepo.Query(ctx, orderID).ExpectOne()
	if err != nil {
		err = fmt.Errorf("getting order: %w", err)
		log.Printf("%s %s", log.Error, err)

		return nil, err
	}

	prd, err := s.ProductRepo.Query(ctx, product.Product{Code: productID}).ExpectOne()
	if err != nil {
		err = fmt.Errorf("getting product: %w", err)
		log.Printf("%s %s", log.Error, err)

		return nil, err
	}

	item := Item{
		Product:  *prd,
		Quantity: qty,
		Price:    prd.Price,
		Total:    prd.Price.Times(qty),
	}

	ord.Items = append(ord.Items, item)
	ord.Total = CalculateTotal(*ord)

	if err := s.OrderRepo.Persist(ctx, ord); err != nil {
		err = fmt.Errorf("persisting order: %w", err)
		log.Printf("%s %s", log.Error, err)

		return nil, err
	}

	return ord, nil
}
