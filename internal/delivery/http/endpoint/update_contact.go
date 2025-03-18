package http

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/niv-e/phonebook-api/internal/application/commands"
	"github.com/niv-e/phonebook-api/internal/application/handlers"
	"github.com/niv-e/phonebook-api/internal/application/model"
	"github.com/niv-e/phonebook-api/internal/domain/repositories"
)

// UpdateContactRequest represents the request payload for updating a contact
type UpdateContactRequest struct {
	ID         uuid.UUID         `json:"id" validate:"required,uuid"`
	FirstName  string            `json:"first_name" validate:"required"`
	LastName   string            `json:"last_name"`
	Phones     []model.PhoneType `json:"phones" validate:"required,dive"`
	Street     string            `json:"street" validate:"required"`
	CityId     uint              `json:"city" validate:"required"`
	PostalCode string            `json:"postal_code"`
	CountryId  uint              `json:"country" validate:"required"`
}

// UpdateContactHttpHandler handles the HTTP request for updating a contact
// @Summary Update a contact
// @Description Update a contact in the phone book
// @Tags contacts
// @Accept  json
// @Produce  json
// @Param contact body UpdateContactRequest true "Contact to update"
// @Success 200 {object} UpdateContactRequest
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to update contact"
// @Router /contacts [put]
func UpdateContactHttpHandler(repo repositories.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateContactRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		cmd := commands.NewUpdateContactCommand(req.ID, req.FirstName, req.LastName, req.Phones, req.Street, req.CityId, req.PostalCode, req.CountryId)
		handler := handlers.NewUpdateContactHandler(repo)
		if err := handler.Handle(cmd); err != nil {
			http.Error(w, "Failed to update contact", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(req)
	}
}
