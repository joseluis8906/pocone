package order

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/pkg/log"
	"github.com/joseluis8906/pocone/pkg/payu"
)

type CreditCard struct {
	Number          string
	ExpirationDate  time.Time
	CVV             string
	CardHoldersName string
	Issuer          string
}

func (s *RpcService) Checkout(ctx context.Context, orderID uuid.UUID, cardInfo CreditCard) error {
	if err := s.PayuGW.Checkout(payu.CheckoutRequest{}); err != nil {
		err := fmt.Errorf("checking out on payment gateway: %w", err)
		log.Printf("%s %s", log.Error, err)

		return err
	}

	return nil
}
