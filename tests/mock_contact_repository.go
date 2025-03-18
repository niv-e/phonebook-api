package tests

import (
	"github.com/google/uuid"
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

func (m *MockContactRepository) Delete(id uuid.UUID) error {
	if m.Err != nil {
		return m.Err
	}
	for i, contact := range m.Contacts {
		if *contact.ID == id {
			m.Contacts = append(m.Contacts[:i], m.Contacts[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockContactRepository) FindByID(id uuid.UUID) (model.ContactType, error) {
	if m.Err != nil {
		return model.ContactType{}, m.Err
	}
	for _, contact := range m.Contacts {
		if *contact.ID == id {
			return contact, nil
		}
	}
	return model.ContactType{}, nil
}

func (m *MockContactRepository) Update(contact model.ContactType) error {
	if m.Err != nil {
		return m.Err
	}
	for i, c := range m.Contacts {
		if *c.ID == *contact.ID {
			m.Contacts[i] = contact
			return nil
		}
	}
	return nil
}

func (m *MockContactRepository) Search(firstName, lastName, fullName, phone string) ([]model.ContactType, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	var results []model.ContactType
	for _, contact := range m.Contacts {
		if (firstName != "" && contact.FirstName == firstName) ||
			(lastName != "" && contact.LastName == lastName) ||
			(fullName != "" && contact.FirstName+" "+contact.LastName == fullName) ||
			(phone != "" && containsPhone(contact.Phones, phone)) {
			results = append(results, contact)
		}
	}
	return results, nil
}

func containsPhone(phones []model.PhoneType, phone string) bool {
	for _, p := range phones {
		if p.Number == phone {
			return true
		}
	}
	return false
}
