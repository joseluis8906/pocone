package customer

import (
	"github.com/joseluis8906/pocone/internal/address"
)

type Customer struct {
	ID        string
	Name      string
	Age       int
	Address   address.Address
	Favorites []string
}
