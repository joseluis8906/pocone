package customer

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/pkg/log"
)

func (s *RpcService) New(ctx context.Context, cus Customer) (*Customer, error) {
	cus.ID = uuid.New()
	if err := s.CustomerRepo.Persist(ctx, &cus); err != nil {
		err = fmt.Errorf("persisting customer: %w", err)
		log.Printf("%s %s", log.Error, err)

		return nil, err
	}

	return &cus, nil
}
