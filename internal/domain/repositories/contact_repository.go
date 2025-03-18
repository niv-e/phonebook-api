package repositories

import (
	"github.com/google/uuid"
	"github.com/niv-e/phonebook-api/internal/application/model"
)

type ContactRepository interface {
	Save(contact model.ContactType) error
	FindPaginated(page, pageSize int) ([]model.ContactType, error)
	Delete(id uuid.UUID) error
}
