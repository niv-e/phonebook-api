package tests

import (
	"github.com/niv-e/phonebook-api/internal/application/model"
)

type MockContactRepository struct {
	Contacts []model.ContactType
	Err      error
}

func (m *MockContactRepository) Save(contact model.ContactType) error {
	if m.Err != nil {
		return m.Err
	}
	m.Contacts = append(m.Contacts, contact)
	return nil
}

func (m *MockContactRepository) FindPaginated(page, pageSize int) ([]model.ContactType, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= len(m.Contacts) {
		return []model.ContactType{}, nil
	}
	if end > len(m.Contacts) {
		end = len(m.Contacts)
	}
	return m.Contacts[start:end], nil
}
