package commands

import (
	"github.com/go-playground/validator/v10"
	"github.com/niv-e/phonebook-api/internal/application/model"
	"github.com/niv-e/phonebook-api/internal/domain"
	"github.com/niv-e/phonebook-api/internal/domain/entity"
)

type AddContactCommand struct {
	FirstName string            `validate:"required"`
	LastName  string            `validate:"required"`
	Phone     model.PhoneType   `validate:"required"`
	Address   model.AddressType `validate:"required"`
}

func NewAddContactCommand(firstName, lastName string, phone model.PhoneType, address model.AddressType) (AddContactCommand, error) {
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

func (c AddContactCommand) ToContact() (entity.ContactEntity, error) {
	contact := entity.ContactEntity{
		ID:        nil,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address: entity.AddressEntity{
			Street:     c.Address.StreetId,
			PostalCode: c.Address.PostalCode,
			City: entity.CityEntity{
				Name: c.Address.CityId,
				Country: entity.CountryEntity{
					Name: c.Address.CountryId,
					// set from country service
					// Alpha2Code:  c.Address.Country.Alpha2Code,
					// Alpha3Code:  c.Address.Country.Alpha3Code,
					// NumericCode: c.Address.Country.NumericCode,
				},
			},
		},
		Phones: []entity.PhoneEntity{
			{
				Number: c.Phone.Value,
				Type:   c.Phone.Type,
			},
		},
	}

	return contact, nil
}
