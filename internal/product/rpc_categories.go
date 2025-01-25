package product

func (s *RpcService) Categories() ([]string, error) {
	return validCategories, nil
}
