package product

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/joseluis8906/pocone/pkg/money"
)

type Product struct {
	Code           string
	Name           string
	Description    string
	Price          money.Money
	Image          string
	Categories     []string
	IsSpecialOffer bool
	Discount       float64
}

func OriginalPrice(p Product) money.Money {
	amount := float64(p.Price.Amount()) / 100
	x := amount / (1 - (p.Discount / 100))
	return money.NewFromFloat(x, p.Code)
}

type addProductTask struct {
	p    Product
	data []byte
	err  error
}

func (a *addProductTask) Unmarshal(r io.Reader) {
	if a.err != nil {
		return
	}

	err := json.NewDecoder(r).Decode(&a.p)
	if err != nil {
		a.err = fmt.Errorf("decoding product: %w", err)
	}
}

func (a *addProductTask) PersistOnDB(ctx context.Context) {
	if a.err != nil {
		return
	}

	if err := persistOnDB(ctx, a.p); err != nil {
		a.err = fmt.Errorf("persisting product: %w", err)
	}
}

func (a *addProductTask) Marshal() {
	if a.err != nil {
		return
	}

	data, err := json.Marshal(a.p)
	if err != nil {
		a.err = fmt.Errorf("marshaling product: %w", err)
	}

	a.data = data
}

func (a *addProductTask) Result() ([]byte, error) {
	return a.data, a.err
}

type getProductTask struct {
	products []Product
	criteria Product
	data     []byte
	err      error
}

func (g *getProductTask) ExtractParams(v url.Values) {
	if g.err != nil {
		return
	}

	g.criteria = Product{
		Code: v.Get("code"),
		Name: v.Get("name"),
	}
}

func (g *getProductTask) SearchOnDB(ctx context.Context) {
	if g.err != nil {
		return
	}

	products, err := queryDB(ctx, g.criteria)
	if err != nil {
		g.err = fmt.Errorf("querying db: %w", err)
		return
	}

	g.products = products
}

func (g *getProductTask) Marshal() {
	if g.err != nil {
		return
	}

	data, err := json.Marshal(g.products)
	if err != nil {
		g.err = fmt.Errorf("marshaling products: %w", err)
		return
	}

	g.data = data
}

func (g *getProductTask) Result() ([]byte, error) {
	return g.data, g.err
}
