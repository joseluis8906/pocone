package product

import (
	"errors"
	"fmt"
	"slices"
)

type validateProduct struct {
	err error
	p   Product
}

func (v *validateProduct) CheckCode() {
	if v.err != nil {
		return
	}

	if v.p.Code == "" {
		v.err = errors.New("empty code")
	}
}

func (v *validateProduct) CheckName() {
	if v.err != nil {
		return
	}

	if v.p.Name == "" {
		v.err = errors.New("empty name")
	}
}

func (v *validateProduct) CheckDescription() {
	if v.err != nil {
		return
	}

	if v.p.Description == "" {
		v.err = errors.New("empty description")
	}
}

func (v *validateProduct) CheckPrice() {
	if v.err != nil {
		return
	}

	if v.p.Price.IsZero() || v.p.Price.IsNegative() {
		v.err = fmt.Errorf("price %s is not allowed", v.p.Price.Display())
	}
}

func (v *validateProduct) CheckImage() {
	if v.err != nil {
		return
	}

	if v.p.Image == "" {
		v.err = errors.New("empty image url")
	}
}

func (v *validateProduct) CheckCategories() {
	if v.err != nil {
		return
	}

	if len(v.p.Categories) == 0 {
		v.err = errors.New("categories are empty")
		return
	}

	for _, c := range v.p.Categories {
		if slices.Contains(validCategories, c) {
			v.err = fmt.Errorf("%s is not an allowed category", c)
			break
		}
	}
}
