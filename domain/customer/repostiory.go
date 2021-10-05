package customer

import (
	"errors"
	"exercise/tavern/aggregate"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("The customer is not present")
	ErrFailedToAddCustomer = errors.New("Failed to create a customer")
	ErrUpdateCustomer      = errors.New("Failed to update a customer")
)

type CustomRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
