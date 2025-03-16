package repositories

import (
	"github.com/niv-e/phonebook-api/internal/application/model"
)

type ContactRepository interface {
	Save(contact model.ContactType) error
	FindPaginated(page, pageSize int) ([]model.ContactType, error)
}
