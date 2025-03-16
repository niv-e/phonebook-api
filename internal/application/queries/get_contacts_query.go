package queries

import (
    "github.com/go-playground/validator/v10"
    "github.com/niv-e/phonebook-api/internal/domain"
)

type GetContactsQuery struct {
    Page int `validate:"min=1"`
}

func NewGetContactsQuery(page int) (GetContactsQuery, error) {
    query := GetContactsQuery{Page: page}
    validate := validator.New()
    if err := validate.Struct(query); err != nil {
        return GetContactsQuery{}, domain.NewInvalidContactError("invalid pagination parameters")
    }
    return query, nil
}