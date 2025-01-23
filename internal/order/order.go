package order

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/joseluis8906/pocone/internal/customer"
	"github.com/joseluis8906/pocone/internal/product"
	"github.com/joseluis8906/pocone/pkg/money"
)

type (
	Order struct {
		ID       uuid.UUID
		Date     time.Time
		Customer customer.Customer
		Items    []Item
		Subtotal money.Money
		Fees     money.Money
		Discount float64
		Total    money.Money
	}

	Item struct {
		Product  product.Product
		Quantity int
		Price    money.Money
		Total    money.Money
	}
)

func encode(od Order) ([]byte, error) {
	data, err := json.Marshal(od)
	if err != nil {
		return data, fmt.Errorf("encoding to json: %w", err)
	}

	return data, nil
}

type CreateTask struct {
	order Order
	data  []byte
	err   error
}

func (c *CreateTask) GenerateID() {
	if c.err != nil {
		return
	}

	c.order.ID = uuid.New()
	c.order.Date = time.Now()
}

func (c *CreateTask) PersistOrder(ctx context.Context) {
	if c.err != nil {
		return
	}

	if err := persistOnDB(ctx, c.order); err != nil {
		c.err = fmt.Errorf("persisting order: %w", err)
	}
}

func (c *CreateTask) MarshalOrder() {
	if c.err != nil {
		return
	}

	c.data, c.err = encode(c.order)
}

func (c *CreateTask) Result() ([]byte, error) {
	return c.data, c.err
}

type GetTask struct {
	order Order
	data  []byte
	err   error
}

func (s *GetTask) ExtractID(r *http.Request) {
	if s.err != nil {
		return
	}

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		s.err = fmt.Errorf("parsing uuid: %w", err)
		return
	}

	s.order.ID = id
}

func (s *GetTask) SearchTheDB(ctx context.Context) {
	if s.err != nil {
		return
	}

	od, err := queryByID(ctx, s.order.ID)
	if err != nil {
		s.err = fmt.Errorf("querying DB: %w", err)
		return
	}

	s.order = od
}

func (s *GetTask) Encode() {
	if s.err != nil {
		return
	}

	s.data, s.err = encode(s.order)
}

func (c *GetTask) Result() ([]byte, error) {
	return c.data, c.err
}
