package commands

import (
	"github.com/google/uuid"
	"github.com/niv-e/phonebook-api/internal/application/model"
)

type UpdateContactCommand struct {
	ID         uuid.UUID
	FirstName  string
	LastName   string
	Phones     []model.PhoneType
	Street     string
	CityId     uint
	PostalCode string
	CountryId  uint
}

func NewUpdateContactCommand(id uuid.UUID, firstName, lastName string, phones []model.PhoneType, street string, cityId uint, postalCode string, countryId uint) UpdateContactCommand {
	return UpdateContactCommand{
		ID:         id,
		FirstName:  firstName,
		LastName:   lastName,
		Phones:     phones,
		Street:     street,
		CityId:     cityId,
		PostalCode: postalCode,
		CountryId:  countryId,
	}
}
