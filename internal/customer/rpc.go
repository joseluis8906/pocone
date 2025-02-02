package customer

import "github.com/joseluis8906/pocone/internal/product"

type RpcService struct {
	CustomerRepo *Repository
	ProductRepo  *product.Repository
}
