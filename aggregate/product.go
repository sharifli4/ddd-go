package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shari4ov/ddd-go.git/entity"
)

var (
	ErrMissingValue = errors.New("missing important value")
)

type Product struct {
	item     *entity.Item
	price    float32
	quantity int
}

func NewProduct(name, description string, price float32) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValue
	}
	return Product{
		item: &entity.Item{
			Name:        name,
			Description: description,
			ID:          uuid.New(),
		},
		price: price,
	}, nil
}
func (p Product) Item() *entity.Item {
	return p.item
}

func (p Product) Price() float32 {
	return p.price
}

func (p Product) Quantity() int {
	return p.quantity
}
