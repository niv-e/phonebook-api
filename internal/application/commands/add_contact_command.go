package commands

import (
	"github.com/niv-e/phonebook-api/internal/domain"
	"github.com/niv-e/phonebook-api/internal/domain/entities"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AddContactCommand struct {
	FirstName string          `validate:"required"`
	LastName  string          `validate:"required"`
	Phone     entities.Phone  `validate:"required"`
	Address   entities.Address `validate:"required"`
}

func NewAddContactCommand(firstName, lastName string, phone entities.Phone, address entities.Address) (AddContactCommand, error) {
	cmd := AddContactCommand{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Address:   address,
	}

	validate := validator.New()
	if err := validate.Struct(cmd); err != nil {
		return AddContactCommand{}, domain.NewInvalidContactError("required fields missing")
	}

	return cmd, nil
}

func (c AddContactCommand) ToContact() (entities.Contact, error) {
	return entities.NewContact(uuid.UUID{}, c.FirstName, c.LastName, c.Phone, c.Address)
}