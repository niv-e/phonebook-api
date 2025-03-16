package entities

import (
	"github.com/niv-e/phonebook-api/internal/domain"
	"github.com/google/uuid"
)

type Contact struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Phone     Phone
	Address   Address
}

func NewContact(id uuid.UUID, firstName, lastName string, phone Phone, address Address) (Contact, error) {
	if firstName == "" {
		return Contact{}, domain.NewInvalidContactError("firstName is required")
	}
	if lastName == "" {
		return Contact{}, domain.NewInvalidContactError("lastName is required")
	}
	return Contact{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Address:   address,
	}, nil
}