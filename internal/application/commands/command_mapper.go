package commands

import (
	"github.com/niv-e/phonebook-api/internal/application/model"
)

func (c AddContactCommand) ToContactType() (model.ContactType, error) {
	contactType := model.ContactType{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address: model.AddressType{
			Street:     c.Address.Street,
			PostalCode: c.Address.PostalCode,
			CityId:     c.Address.CityId,
			CountryId:  c.Address.CountryId,
		},
		Phones: c.Phones,
	}

	return contactType, nil
}
