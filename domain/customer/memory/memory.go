// Package memory is an in-memory implementation of Customer repository
package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shari4ov/ddd-go.git/aggregate"
	"github.com/shari4ov/ddd-go.git/domain/customer"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New(customers map[uuid.UUID]aggregate.Customer) *MemoryRepository {
	return &MemoryRepository{
		customers: customers,
	}
}
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}
func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// Make sure customer is already in repo
	if _, ok := mr.customers[c.Person.ID]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.Person.ID] = c
	mr.Unlock()
	return nil
}
func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.Person.ID]; !ok {
		return fmt.Errorf("customer does not exists: %w", customer.ErrUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.Person.ID] = c
	mr.Unlock()
	return nil
}
