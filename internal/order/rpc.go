package order

import (
	"github.com/joseluis8906/pocone/internal/customer"
	"github.com/joseluis8906/pocone/internal/product"
	"github.com/spf13/viper"
)

type RpcService struct {
	Conf         *viper.Viper
	OrderRepo    *Repository
	ProductRepo  *product.Repository
	CustomerRepo *customer.Repository
}
