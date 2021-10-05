package tests

import (
	"exercise/tavern/aggregate"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty Name Validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "Valid Name",
			name:        "Karim",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		//RUN TEST
		t.Run(tc.test, func(t *testing.T) {
			//CREATE A NEW CUSTOMER
			_, err := aggregate.NewCustomer(tc.name)
			// CHECK IF ERROR MATCH EXPECTED ERROR
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
