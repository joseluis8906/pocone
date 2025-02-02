package product

import (
	"context"
	"fmt"

	"github.com/joseluis8906/pocone/pkg/log"
)

func (s *RpcService) Add(ctx context.Context, req *Product) error {
	task := validateProduct{p: *req}
	task.CheckCode()
	task.CheckName()
	task.CheckDescription()
	task.CheckCategories()
	task.CheckImage()
	task.CheckPrice()

	if task.err != nil {
		err := fmt.Errorf("validating product: %w", task.err)
		log.Printf("%s %v", log.Error, err)
		return err
	}

	if err := s.Repository.Persist(context.Background(), req); err != nil {
		err := fmt.Errorf("persisting product: %w", err)
		log.Printf("%s %v", log.Error, err)
		return err
	}

	return nil
}
