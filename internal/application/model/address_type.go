package model

import "github.com/niv-e/phonebook-api/internal/domain"

type AddressType struct {
	Street     string
	CityId     uint
	CityName  string
	PostalCode string
	CountryName string
	CountryId  uint
}

func NewAddress(street, postalCode string, cityId, countryId uint) (AddressType, error) {
	if street == "" {
		return AddressType{}, domain.NewInvalidAddressError("street is required")
	}
	if cityId == 0 {
		return AddressType{}, domain.NewInvalidAddressError("city is required")
	}

	return AddressType{
		Street:     street,
		CityId:     cityId,
		PostalCode: postalCode,
		CountryId:  countryId,
	}, nil
}
