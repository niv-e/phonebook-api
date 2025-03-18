package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/niv-e/phonebook-api/internal/delivery/http/endpoint"
	"github.com/stretchr/testify/assert"
)

func TestAddContactHttpHandler_Success(t *testing.T) {
	mockRepo := &MockContactRepository{}

	reqBody := map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"phones": []map[string]interface{}{
			{"number": "+12025550123", "type": "mobile"},
		},
		"street":      "123 Main St",
		"city":        1,
		"postal_code": "10001",
		"country":     1,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/contacts", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.AddContactHttpHandler(mockRepo))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Len(t, mockRepo.Contacts, 1)
	assert.Equal(t, "John", mockRepo.Contacts[0].FirstName)
	assert.Equal(t, "Doe", mockRepo.Contacts[0].LastName)
}

func TestAddContactHttpHandler_PayloadMissingRequiredFields_ShouldReturnInvalidPayload(t *testing.T) {
	reqBody := map[string]interface{}{
		"first_name": "John",
		// Missing required fields
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/contacts", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.AddContactHttpHandler(nil))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "INVALID_CONTACT: invalid contact: required fields missing")
}

func TestAddContactHttpHandler_PhoneIsNotInFormatE164_ShouldReturnInvalidPhoneNumber(t *testing.T) {
	mockRepo := &MockContactRepository{}

	reqBody := map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"phones": []map[string]interface{}{
			{"number": "123-456-7890", "type": "mobile"}, // Invalid phone number format
		},
		"street":      "123 Main St",
		"city":        1,
		"postal_code": "10001",
		"country":     1,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/contacts", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.AddContactHttpHandler(mockRepo))
	// assert.Panics(t, func() {
	// 	handler.ServeHTTP(rr, req)
	// })
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "INVALID_CONTACT")
}

func TestAddContactHttpHandler_ValidPhoneNumber_ShouldReturnSuccess(t *testing.T) {
	mockRepo := &MockContactRepository{}

	reqBody := map[string]interface{}{
		"first_name": "Jane",
		"last_name":  "Doe",
		"phones": []map[string]interface{}{
			{"number": "+1234567890", "type": "mobile"}, // Valid phone number format
		},
		"street":      "456 Elm St",
		"city":        2,
		"postal_code": "20002",
		"country":     2,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/contacts", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.AddContactHttpHandler(mockRepo))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Len(t, mockRepo.Contacts, 1)
	assert.Equal(t, "Jane", mockRepo.Contacts[0].FirstName)
	assert.Equal(t, "Doe", mockRepo.Contacts[0].LastName)
}

func TestAddContactHttpHandler_InternalServerError_ShouldReturnInternalServerError(t *testing.T) {
	mockRepo := &MockContactRepository{
		Err: assert.AnError,
	}

	reqBody := map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"phones": []map[string]interface{}{
			{"number": "+12025550123", "type": "mobile"},
		},
		"street":      "123 Main St",
		"city":        1,
		"postal_code": "10001",
		"country":     1,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/contacts", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.AddContactHttpHandler(mockRepo))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
