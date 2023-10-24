package services

import (
	"github.com/google/uuid"
	"github.com/shari4ov/ddd-go.git/aggregate"
	"github.com/shari4ov/ddd-go.git/domain/customer"
	"github.com/shari4ov/ddd-go.git/domain/customer/memory"
	"github.com/shari4ov/ddd-go.git/domain/product"
	prodmem "github.com/shari4ov/ddd-go.git/domain/product/memory"
	"log"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range cfgs {
		if err := cfg(os); err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()
		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}
func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) error {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return err
	}
	var products []aggregate.Product
	var totalPrice float32
	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return err
		}
		products = append(products, p)
		totalPrice += p.Price()
	}
	log.Printf("Customer: %s has ordered %d products", c.Person.ID, len(products))
	return nil
}
