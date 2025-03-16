package model

import "github.com/niv-e/phonebook-api/internal/domain"

type AddressType struct {
	Street     string
	CityId     string
	PostalCode string
	CountryId  string
}

func NewAddress(streetId, cityId, postalCode, countryId string) (AddressType, error) {
	if streetId == "" {
		return AddressType{}, domain.NewInvalidAddressError("street is required")
	}
	if cityId == "" {
		return AddressType{}, domain.NewInvalidAddressError("city is required")
	}

	return AddressType{
		Street:     streetId,
		CityId:     cityId,
		PostalCode: postalCode,
		CountryId:  countryId,
	}, nil
}
