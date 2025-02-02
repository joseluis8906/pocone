package product

import "context"

func (s *RpcService) Categories(ctx context.Context) ([]string, error) {
	return validCategories, nil
}
