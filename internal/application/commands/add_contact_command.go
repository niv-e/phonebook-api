package commands

import (
	"github.com/go-playground/validator/v10"
	"github.com/niv-e/phonebook-api/internal/application/model"
	"github.com/niv-e/phonebook-api/internal/domain"
)

type AddContactCommand struct {
	FirstName string            `validate:"required"`
	LastName  string            `validate:"required"`
	Phones    []model.PhoneType `validate:"required"`
	Address   model.AddressType `validate:"required"`
}

func NewAddContactCommand(firstName, lastName string, phones []model.PhoneType, address model.AddressType) (AddContactCommand, error) {
	cmd := AddContactCommand{
		FirstName: firstName,
		LastName:  lastName,
		Phones:    phones,
		Address:   address,
	}

	validate := validator.New()
	if err := validate.Struct(cmd); err != nil {
		return AddContactCommand{}, domain.NewInvalidContactError("required fields missing")
	}

	return cmd, nil
}
