package aggregate

import (
	"errors"
	"exercise/tavern/entity"
	"exercise/tavern/valueobject"

	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer need to have a valid person")
)

type Customer struct {
	person *entity.Person

	products []*entity.Item

	transaction []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	//Create a new person and generate the ID
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:      person,
		products:    make([]*entity.Item, 0),
		transaction: make([]valueobject.Transaction, 0),
	}, nil
}

// GET ID
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SET ID
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

// UPDATE NAME
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

// GET NAME
func (c *Customer) GetName() string {
	return c.person.Name
}
