package order

import (
	"context"

	"github.com/google/uuid"
)

func (s *RpcService) Checkout(ctx context.Context, orderID uuid.UUID) error {
	return nil
}
