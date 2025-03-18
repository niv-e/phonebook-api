package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/niv-e/phonebook-api/internal/application/model"
	api "github.com/niv-e/phonebook-api/internal/delivery/http/endpoint"
	"github.com/stretchr/testify/assert"
)

func TestUpdateContactHttpHandler_Success(t *testing.T) {
	mockRepo := &MockContactRepository{}
	id := uuid.New()
	mockRepo.Contacts = append(mockRepo.Contacts, model.ContactType{
		ID:        &id,
		FirstName: "Initial",
		LastName:  "Contact",
		Phones:    []model.PhoneType{{Number: "+12025550123", Type: "mobile"}},
		Address: model.AddressType{
			Street:     "123 Main St",
			CityId:     1,
			PostalCode: "10001",
			CountryId:  1,
		},
	})

	reqBody := map[string]interface{}{
		"id":         id.String(),
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

	req, err := http.NewRequest("PUT", "/contacts", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.UpdateContactHttpHandler(mockRepo))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Len(t, mockRepo.Contacts, 1)
	assert.Equal(t, "John", mockRepo.Contacts[0].FirstName)
	assert.Equal(t, "Doe", mockRepo.Contacts[0].LastName)
}

func TestUpdateContactHttpHandler_InvalidID_ShouldReturnInvalidPayload(t *testing.T) {
	reqBody := map[string]interface{}{
		"id":         "invalid-uuid",
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

	req, err := http.NewRequest("PUT", "/contacts", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.UpdateContactHttpHandler(nil))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Invalid request payload\n", rr.Body.String())
}

func TestUpdateContactHttpHandler_InternalServerError_ShouldReturnInternalServerError(t *testing.T) {
	mockRepo := &MockContactRepository{
		Err: assert.AnError,
	}
	id := uuid.New()

	reqBody := map[string]interface{}{
		"id":         id.String(),
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

	req, err := http.NewRequest("PUT", "/contacts", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.UpdateContactHttpHandler(mockRepo))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
