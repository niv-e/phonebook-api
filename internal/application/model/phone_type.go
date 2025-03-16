package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type PhoneType struct {
	Number string
	Type   string
}

func NewPhone(phoneType string, value string) (PhoneType, error) {

	validate := validator.New()
	if err := validate.Var(value, "e164"); err != nil {
		return PhoneType{}, err
	}
	return PhoneType{Type: phoneType, Number: value}, nil
}

func (p PhoneType) String() string {
	return fmt.Sprintf("%s: %s", p.Type, p.Number)
}
