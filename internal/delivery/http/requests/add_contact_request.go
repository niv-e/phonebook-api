package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/niv-e/phonebook-api/internal/application/commands"
	"github.com/niv-e/phonebook-api/internal/application/dto"
	"github.com/niv-e/phonebook-api/internal/application/model"
	"github.com/niv-e/phonebook-api/internal/domain"
)

type AddContactRequest struct {
	FirstName  string       `json:"first_name" validate:"required"`
	LastName   string       `json:"last_name"`
	Phone      dto.PhoneDTO `json:"phone" validate:"required"`
	StreetId   string       `json:"street" validate:"required"`
	CityId     string       `json:"city" validate:"required"`
	PostalCode string       `json:"postal_code"`
	CountryId  string       `json:"country" validate:"required"`
}

func (r AddContactRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return domain.NewInvalidContactError("required fields missing")
	}
	return nil
}

func (r AddContactRequest) ToCommand() (commands.AddContactCommand, error) {
	phone, err := model.NewPhone(r.Phone.Type, r.Phone.Number)
	if err != nil {
		return commands.AddContactCommand{}, err
	}

	address, err := model.NewAddress(r.StreetId, r.CityId, r.PostalCode, r.CountryId)
	if err != nil {
		return commands.AddContactCommand{}, err
	}

	return commands.NewAddContactCommand(r.FirstName, r.LastName, phone, address)
}
