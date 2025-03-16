package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/niv-e/phonebook-api/internal/application/commands"
	"github.com/niv-e/phonebook-api/internal/domain"
	"github.com/niv-e/phonebook-api/internal/domain/entities"
)

type AddContactRequest struct {
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone" validate:"required"`
	Street     string `json:"street" validate:"required"`
	City       string `json:"city" validate:"required"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country" validate:"required"`
}

func (r AddContactRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return domain.NewInvalidContactError("required fields missing")
	}

	// Validate Phone as E.164
	if _, err := entities.NewPhone(r.Phone); err != nil {
		return err
	}

	// Validate Address fields
	if _, err := entities.NewAddress(r.Street, r.City, r.PostalCode, r.Country); err != nil {
		return err
	}

	return nil
}

func (r AddContactRequest) ToCommand() (commands.AddContactCommand, error) {
	phone, err := entities.NewPhone(r.Phone)
	if err != nil {
		return commands.AddContactCommand{}, err
	}

	address, err := entities.NewAddress(r.Street, r.City, r.PostalCode, r.Country)
	if err != nil {
		return commands.AddContactCommand{}, err
	}

	return commands.NewAddContactCommand(r.FirstName, r.LastName, phone, address)
}
