package order

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/internal/customer"
	"github.com/joseluis8906/pocone/pkg/log"
	"github.com/joseluis8906/pocone/pkg/money"
)

func (s *RpcService) New(ctx context.Context, customerID uuid.UUID) (*Order, error) {
	cus, err := s.CustomerRepo.Query(ctx, customer.Customer{ID: customerID}).ExpectOne()
	if err != nil {
		err = fmt.Errorf("getting customer: %v", err)
		log.Printf("%s %s", log.Error, err)

		return nil, err
	}

	ord := Order{
		ID:       uuid.New(),
		Date:     time.Now(),
		Customer: *cus,
		Total:    money.New(0, s.Conf.GetString("currency")),
	}

	if err := s.OrderRepo.Persist(ctx, &ord); err != nil {
		err = fmt.Errorf("persisting order: %v", err)
		log.Printf("%s %s", log.Error, err)

		return nil, err
	}

	return &ord, nil
}
