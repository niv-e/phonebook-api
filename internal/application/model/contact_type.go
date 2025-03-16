package model

import (
	"github.com/google/uuid"
	"github.com/niv-e/phonebook-api/internal/domain"
)

type ContactType struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Phone     PhoneType
	Address   AddressType
}

func NewContact(id uuid.UUID, firstName, lastName string, phone PhoneType, address AddressType) (ContactType, error) {
	if firstName == "" {
		return ContactType{}, domain.NewInvalidContactError("firstName is required")
	}
	if lastName == "" {
		return ContactType{}, domain.NewInvalidContactError("lastName is required")
	}
	return ContactType{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Address:   address,
	}, nil
}
