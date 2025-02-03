package customer

import (
	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/internal/address"
)

type Customer struct {
	ID        uuid.UUID
	Name      string
	Age       int
	Address   []address.Address
	Favorites []string
	Scores    []Score
}

type Score struct {
	Product string
	Score   float32
}
