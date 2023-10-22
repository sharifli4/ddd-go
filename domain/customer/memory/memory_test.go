package memory

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shari4ov/ddd-go.git/aggregate"
	"github.com/shari4ov/ddd-go.git/domain/customer"
	"testing"
)

func TestMemoryRepository_Get(t *testing.T) {
	type testCase struct {
		name          string
		id            uuid.UUID
		expectedError error
	}
	cust, err := aggregate.NewCustomer("penny")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.Person.ID
	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}
	testCases := []testCase{
		{
			name:          "no customer by id",
			id:            uuid.New(),
			expectedError: customer.ErrCustomerNotFound,
		},
		{
			name:          "customer found",
			id:            id,
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error: %v, got %v ", tc.expectedError, err)
			}
		})
	}
}
