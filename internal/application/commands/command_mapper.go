package commands

import (
	"github.com/niv-e/phonebook-api/internal/application/model"
	"github.com/niv-e/phonebook-api/internal/domain/entity"
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
func (c AddContactCommand) ToContactEntity() (entity.ContactEntity, error) {
	phones := make([]entity.PhoneEntity, len(c.Phones))
	for i, phone := range c.Phones {
		phones[i] = entity.PhoneEntity{
			Number: phone.Number,
			Type:   phone.Type,
		}
	}

	contact := entity.ContactEntity{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address: entity.AddressEntity{
			Street:     c.Address.Street,
			PostalCode: c.Address.PostalCode,
			CityID:     c.Address.CityId,
		},
		Phones: phones,
	}

	return contact, nil
}
