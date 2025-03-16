package repositories

import (
	"github.com/niv-e/phonebook-api/internal/domain/entities"
)

type ContactRepository interface {
	Save(contact entities.Contact) error
	// Future methods: FindByID, FindAll, etc.
}
