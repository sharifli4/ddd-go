package aggregate_test

import (
	"errors"
	"github.com/shari4ov/ddd-go.git/aggregate"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test          string
		name          string
		expectedError error
	}
	testCases := []testCase{
		{
			test:          "Empty name validation",
			name:          "",
			expectedError: aggregate.ErrInvalidPerson,
		},
		{
			test:          "Valid name",
			name:          "Kenan",
			expectedError: nil,
		},
	}
	for _, v := range testCases {
		t.Run(v.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(v.name)
			if !errors.Is(err, v.expectedError) {
				t.Errorf("expected error %v, got %v", v.expectedError, err)
			}
		})
	}
}
