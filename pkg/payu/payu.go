package payu

import "net/http"

type (
	Gateway struct{}

	Config struct{}
)

type CheckoutRequest struct {
	client *http.Client
}

func New(config Config) *Gateway {
	return &Gateway{}
}

func (g *Gateway) Checkout(req CheckoutRequest) error {
	return nil
}
