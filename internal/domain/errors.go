package domain

import "fmt"

type DomainError struct {
	Code    string // e.g., "INVALID_PHONE"
	Message string
}

func (e DomainError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Error codes
const (
	ErrInvalidPhoneCode   = "INVALID_PHONE"
	ErrInvalidAddressCode = "INVALID_ADDRESS"
	ErrInvalidContactCode = "INVALID_CONTACT"
)

// Factory methods for specific domain errors
func NewInvalidPhoneError(phone string) DomainError {
	return DomainError{
		Code:    ErrInvalidPhoneCode,
		Message: fmt.Sprintf("phone '%s' must follow E.164 format (e.g., +12025550123)", phone),
	}
}

func NewInvalidAddressError(details string) DomainError {
	return DomainError{
		Code:    ErrInvalidAddressCode,
		Message: fmt.Sprintf("invalid address: %s", details),
	}
}

func NewInvalidContactError(details string) DomainError {
	return DomainError{
		Code:    ErrInvalidContactCode,
		Message: fmt.Sprintf("invalid contact: %s", details),
	}
}
