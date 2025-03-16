package tests

import (
	"testing"

	"github.com/niv-e/phonebook-api/internal/application/commands"
	"github.com/niv-e/phonebook-api/internal/application/handlers"
	"github.com/niv-e/phonebook-api/internal/domain"
	"github.com/niv-e/phonebook-api/internal/domain/entities"
)

type mockContactRepository struct {
	contacts []entities.Contact
}

func (m *mockContactRepository) Save(contact entities.Contact) error {
	m.contacts = append(m.contacts, contact)
	return nil
}

func TestAddContactHandler_HandleCommandWithAllRequiredFields_ShouldAddNewContact(t *testing.T) {
	repo := &mockContactRepository{}
	handler := handlers.NewAddContactHandler(repo)

	phone, _ := entities.NewPhone("+12025550123")
	address, _ := entities.NewAddress("123 Main St", "New York", "10001", "USA")
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
	if savedContact.Phone.String() != "+12025550123" {
		t.Errorf("Expected Phone '+12025550123', got %s", savedContact.Phone.String())
	}
}

func TestAddContactHandler_HandleCommandWithMissingFields_ShouldReturnError(t *testing.T) {
	repo := &mockContactRepository{}
	handler := handlers.NewAddContactHandler(repo)

	phone, _ := entities.NewPhone("+12025550123")
	address, _ := entities.NewAddress("123 Main St", "New York", "10001", "USA")
	cmd, err := commands.NewAddContactCommand("", "Doe", phone, address)
	if err == nil {
		t.Fatalf("Expected error for missing FirstName, got nil")
	}

	err = handler.Handle(cmd)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestAddContactCommand_Validation_ShouldReturnErrorForInvalidFields(t *testing.T) {
	phone, _ := entities.NewPhone("+12025550123")
	address, _ := entities.NewAddress("123 Main St", "New York", "10001", "USA")

	_, err := commands.NewAddContactCommand("", "Doe", phone, address)
	if err == nil {
		t.Errorf("Expected error for empty FirstName, got nil")
	} else if de, ok := err.(domain.DomainError); !ok || de.Code != domain.ErrInvalidContactCode {
		t.Errorf("Expected DomainError with code %s, got %v", domain.ErrInvalidContactCode, err)
	}
}
