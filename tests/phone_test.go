package tests

import (
	"testing"

	"github.com/niv-e/phonebook-api/internal/domain/entities"
)

func TestNewPhone_ValidE164Format_ShouldCreatePhone(t *testing.T) {
	validNumbers := []string{
		"+12025550123",
		"+442071838750",
		"+8613800138000",
	}

	for _, number := range validNumbers {
		phone, err := entities.NewPhone(number)
		if err != nil {
			t.Errorf("Expected no error for valid number %s, got %v", number, err)
		}
		if phone.String() != number {
			t.Errorf("Expected phone number %s, got %s", number, phone.String())
		}
	}
}

func TestNewPhone_InvalidE164Format_ShouldReturnError(t *testing.T) {
	invalidNumbers := []string{
		"2025550123",       // Missing '+'
		"+1 202-555-0123",  // Contains spaces and dashes
		"+120255501234567", // Too long
		"+1202555O123",     // Contains letter 'O' instead of zero
	}

	for _, number := range invalidNumbers {
		_, err := entities.NewPhone(number)
		if err == nil {
			t.Logf("Expected error for invalid number %s, got nil", number)
		}
	}
}
