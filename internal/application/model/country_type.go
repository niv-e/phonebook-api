package model

import "github.com/niv-e/phonebook-api/internal/domain"

type CountryType struct {
    Name        string
    Alpha2Code  string
    Alpha3Code  string
    NumericCode string
}

func NewCountry(name, alpha2Code, alpha3Code, numericCode string) (CountryType, error) {
    if name == "" {
        return CountryType{}, domain.NewInvalidAddressError("country name is required")
    }
    if alpha2Code == "" || len(alpha2Code) != 2 {
        return CountryType{}, domain.NewInvalidAddressError("valid alpha-2 code is required")
    }
    if alpha3Code == "" || len(alpha3Code) != 3 {
        return CountryType{}, domain.NewInvalidAddressError("valid alpha-3 code is required")
    }
    if numericCode == "" || len(numericCode) != 3 {
        return CountryType{}, domain.NewInvalidAddressError("valid numeric code is required")
    }
    return CountryType{
        Name:        name,
        Alpha2Code:  alpha2Code,
        Alpha3Code:  alpha3Code,
        NumericCode: numericCode,
    }, nil
}