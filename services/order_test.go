package services

import (
	"github.com/google/uuid"
	"github.com/shari4ov/ddd-go.git/aggregate"
	"testing"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy beverage", 11)
	if err != nil {
		t.Fatal(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts", "Snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}
	wine, err := aggregate.NewProduct("Wine", "Some drink", 25)
	if err != nil {
		t.Fatal(err)
	}
	return []aggregate.Product{
		beer, peenuts, wine,
	}
}
func TestNewOrderService(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products))
	if err != nil {
		t.Fatal(err)
	}
	cust, err := aggregate.NewCustomer("Kenan")
	if err != nil {
		t.Error(err)
	}
	if err := os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}
	order := []uuid.UUID{
		products[0].Item().ID,
	}
	if err := os.CreateOrder(cust.Person.ID, order); err != nil {
		t.Error(err)
	}
}
