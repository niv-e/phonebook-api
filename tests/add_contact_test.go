package tests

import (
	"testing"

	"github.com/niv-e/phonebook-api/internal/application/commands"
	"github.com/niv-e/phonebook-api/internal/application/handlers"
	"github.com/niv-e/phonebook-api/internal/application/model"
	"github.com/niv-e/phonebook-api/internal/domain"
	"github.com/niv-e/phonebook-api/internal/domain/entity"
)

type mockContactRepository struct {
	contacts []entity.ContactEntity
}

func (m *mockContactRepository) Save(contact entity.ContactEntity) error {
	m.contacts = append(m.contacts, contact)
	return nil
}

func TestAddContactHandler_HandleCommandWithAllRequiredFields_ShouldAddNewContact(t *testing.T) {
	repo := &mockContactRepository{}
	handler := handlers.NewAddContactHandler(repo)

	phone, _ := model.NewPhone("mobile", "+12025550123")
	address, _ := model.NewAddress("123", "NY", "10001", "USA")
	cmd, err := commands.NewAddContactCommand("John", "Doe", phone, address)
	if err != nil {
		t.Fatalf("Failed to create command: %v", err)
	}

	err = handler.Handle(cmd)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(repo.contacts) != 1 {
		t.Errorf("Expected 1 contact in repo, got %d", len(repo.contacts))
	}

	savedContact := repo.contacts[0]
	if savedContact.FirstName != "John" {
		t.Errorf("Expected FirstName 'John', got %s", savedContact.FirstName)
	}
	if savedContact.Phones[0].Number != "+12025550123" {
		t.Errorf("Expected Phone '+12025550123', got %s", savedContact.Phones[0].Number)
	}
}

func TestAddContactHandler_HandleCommandWithMissingFields_ShouldReturnError(t *testing.T) {
	repo := &mockContactRepository{}
	handler := handlers.NewAddContactHandler(repo)

	phone, _ := model.NewPhone("mobile", "+12025550123")
	address, _ := model.NewAddress("123", "NY", "10001", "USA")
	cmd, err := commands.NewAddContactCommand("", "Doe", phone, address)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if domainErr, ok := err.(domain.DomainError); !ok || domainErr.Code != domain.ErrInvalidContactCode {
		t.Errorf("Expected domain error with code %v, got %v", domain.ErrInvalidContactCode, err)
	}

	_ = handler.Handle(cmd)
}
