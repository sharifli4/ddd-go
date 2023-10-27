package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/shari4ov/ddd-go.git/aggregate"
	"testing"
)

func TestTavern_Order(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}
	cust, err := aggregate.NewCustomer("Kenan")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}
	order := []uuid.UUID{
		products[0].Item().ID,
	}
	if err := tavern.Order(cust.Person.ID, order); err != nil {
		t.Fatal(err)
	}

}
