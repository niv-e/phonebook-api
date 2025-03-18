package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/niv-e/phonebook-api/internal/application/model"
	api "github.com/niv-e/phonebook-api/internal/delivery/http/endpoint"
	"github.com/stretchr/testify/assert"
)

func TestGetContactsHttpHandler_Success(t *testing.T) {
	mockRepo := &MockContactRepository{
		Contacts: []model.ContactType{
			{
				FirstName: "John",
				LastName:  "Doe",
				Address: model.AddressType{
					Street:      "123 Main St",
					CityId:      1,
					CityName:    "New York",
					CountryId:   1,
					CountryName: "United States",
				},
				Phones: []model.PhoneType{
					{Number: "+12025550123", Type: "mobile"},
				},
			},
		},
	}

	req, err := http.NewRequest("GET", "/contacts?page=1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetContactsHttpHandler(mockRepo))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var contacts []model.ContactType
	err = json.NewDecoder(rr.Body).Decode(&contacts)
	assert.NoError(t, err)
	assert.Len(t, contacts, 1)
	assert.Equal(t, "John", contacts[0].FirstName)
	assert.Equal(t, "Doe", contacts[0].LastName)
}

func TestGetContactsHttpHandler_InvalidPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/contacts?page=0", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetContactsHttpHandler(nil))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Invalid page number\n", rr.Body.String())
}

func TestGetContactsHttpHandler_InternalServerError(t *testing.T) {
	mockRepo := &MockContactRepository{
		Err: assert.AnError,
	}

	req, err := http.NewRequest("GET", "/contacts?page=1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetContactsHttpHandler(mockRepo))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
