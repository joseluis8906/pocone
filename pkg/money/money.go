package money

import (
	"encoding/json"

	"github.com/Rhymond/go-money"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Money struct {
	money.Money
}

func New(amount int64, code string) Money {
	return Money{*money.New(amount, code)}
}

func COP(integer int64, decimal int64) Money {
	return Money{*money.New(integer*100+decimal, money.COP)}
}

func USD(integer int64, decimal int64) Money {
	return Money{*money.New(integer*100+decimal, money.USD)}
}

func NewFromFloat(amount float64, code string) Money {
	return Money{*money.NewFromFloat(amount, code)}
}

func (m Money) MarshalBSON() ([]byte, error) {
	v := struct {
		Amount   int64
		Currency string
	}{
		Amount:   m.Amount(),
		Currency: m.Currency().Code,
	}

	return bson.Marshal(v)
}

func toFloat(m Money) float64 {
	return float64(m.Amount()) / 100.0
}

func (m *Money) UnmarshalJSON(data []byte) error {
	var raw struct {
		Amount   float64
		Currency string
	}

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	v := money.NewFromFloat(raw.Amount, raw.Currency)
	m.Money = *v

	return nil
}

func (m Money) MarshalJSON() ([]byte, error) {
	v := struct {
		Amount   float64
		Currency string
	}{
		Amount:   toFloat(m),
		Currency: m.Currency().Code,
	}

	return json.Marshal(v)
}

func (m *Money) UnmarshalBSON(data []byte) error {
	var raw struct {
		Amount   int64
		Currency string
	}

	err := bson.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	v := money.New(raw.Amount, raw.Currency)
	m.Money = *v

	return nil
}
