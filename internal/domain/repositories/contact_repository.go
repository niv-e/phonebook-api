package repositories

import (
	"github.com/niv-e/phonebook-api/internal/domain/entity"
)

type ContactRepository interface {
	Save(contact entity.ContactEntity) error
	// Future methods: FindByID, FindAll, etc.
}
