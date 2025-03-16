package entities

import (
	"github.com/go-playground/validator/v10"
)

type Phone struct {
	value string
}

func NewPhone(value string) (Phone, error) {

	validate := validator.New()
	if err := validate.Var(value, "e164"); err != nil {
		return Phone{}, err
	}
	return Phone{value: value}, nil
}

func (p Phone) String() string {
	return p.value
}
