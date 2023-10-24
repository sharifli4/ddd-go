package product

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shari4ov/ddd-go.git/aggregate"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductExists   = errors.New("product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(uuid.UUID) (aggregate.Product, error)
	Add(aggregate.Product) error
	Update(aggregate.Product) error
	Delete(uuid.UUID) error
}
