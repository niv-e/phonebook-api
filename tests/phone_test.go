package tests

import (
	"testing"

	"github.com/niv-e/phonebook-api/internal/application/dto"
	"github.com/niv-e/phonebook-api/internal/application/model"
)

func TestNewPhone_ValidE164Format_ShouldCreatePhone(t *testing.T) {
	validNumbers := []dto.PhoneDTO{
		{Type: "mobile", Number: "+12025550123"},
		{Type: "mobile", Number: "+442071838750"},
		{Type: "mobile", Number: "+8613800138000"},
	}

	for _, number := range validNumbers {
		_, err := model.NewPhone(number.Type, number.Number)
		if err != nil {
			t.Errorf("Expected no error for valid number %s, got %v", number, err)
		}
	}
}

func TestNewPhone_InvalidE164Format_ShouldReturnError(t *testing.T) {
	invalidNumbers := []dto.PhoneDTO{
		{Type: "mobile", Number: "2025550123"},       // Missing '+'
		{Type: "mobile", Number: "+1 202-555-0123"},  // Contains spaces and dashes
		{Type: "mobile", Number: "+120255501234567"}, // Too long
		{Type: "mobile", Number: "+1202555O123"},     // Contains letter 'O' instead of zero
	}

	for _, number := range invalidNumbers {
		_, err := model.NewPhone(number.Type, number.Number)
		if err == nil {
			t.Logf("Expected error for invalid number %s, got nil", number)
		}
	}
}
