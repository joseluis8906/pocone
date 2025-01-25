package customer

import "github.com/joseluis8906/pocone/internal/product"

type RpcService struct {
	Repository  *Repository
	ProductRepo *product.Repository
}
