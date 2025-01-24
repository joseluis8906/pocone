package product

import (
	"context"
	"fmt"
)

func (s *RpcService) Add(req *Product) error {
	task := validateProduct{p: *req}
	task.CheckCode()
	task.CheckName()
	task.CheckDescription()
	task.CheckCategories()
	task.CheckImage()
	task.CheckPrice()

	if task.err != nil {
		return fmt.Errorf("validating product: %w", task.err)
	}

	if err := s.Repository.persist(context.Background(), *req); err != nil {
		return fmt.Errorf("persisting product: %w", err)
	}

	return nil
}
