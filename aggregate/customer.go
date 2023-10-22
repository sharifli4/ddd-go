// Package aggregate holds our aggregates that combines many entities into a full object
package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shari4ov/ddd-go.git/entity"
	"github.com/shari4ov/ddd-go.git/valueobject"
)

type Customer struct {
	// Person is the root entity of customer
	// which means person.ID is the main identifier for the customer
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

var (
	ErrInvalidPerson = errors.New("invalid customer name")
)

// NewCustomer is a factory to create a new customer aggregate
// it will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}
	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}
