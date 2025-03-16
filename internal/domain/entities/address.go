package entities

import "github.com/niv-e/phonebook-api/internal/domain"

type Address struct {
	Street     string
	City       string
	PostalCode string
	Country    string
}

func NewAddress(street, city, postalCode, country string) (Address, error) {
	if street == "" {
		return Address{}, domain.NewInvalidAddressError("street is required")
	}
	if city == "" {
		return Address{}, domain.NewInvalidAddressError("city is required")
	}
	if country == "" {
		return Address{}, domain.NewInvalidAddressError("country is required")
	}
	return Address{
		Street:     street,
		City:       city,
		PostalCode: postalCode,
		Country:    country,
	}, nil
}