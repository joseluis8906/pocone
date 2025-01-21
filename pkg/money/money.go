package money

import (
	"github.com/Rhymond/go-money"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Money struct {
	money.Money
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
