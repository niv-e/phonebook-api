package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/niv-e/phonebook-api/internal/application/commands"
	"github.com/niv-e/phonebook-api/internal/application/handlers"
	"github.com/niv-e/phonebook-api/internal/application/model"
	"github.com/niv-e/phonebook-api/internal/domain"
	"github.com/niv-e/phonebook-api/internal/domain/repositories"
)

// AddContactRequest represents the request payload for adding a contact
//
//	type AddContactRequest struct {
//		FirstName  string            `json:"first_name" validate:"required"`
//		LastName   string            `json:"last_name"`
//		Phones     []model.PhoneType `json:"phones" validate:"required,dive"`
//		Street     string            `json:"street" validate:"required"`
//		CityId     uint              `json:"city" validate:"required"`
//		PostalCode string            `json:"postal_code"`
//		CountryId  uint              `json:"country" validate:"required"`
//	}
type AddContactRequest struct {
	FirstName  string            `json:"first_name" validate:"required"`
	LastName   string            `json:"last_name"`
	Phones     []model.PhoneType `json:"phones" validate:"required,dive"`
	Street     string            `json:"street" validate:"required"`
	CityId     uint              `json:"city" validate:"required"`
	PostalCode string            `json:"postal_code"`
	CountryId  uint              `json:"country" validate:"required"`
}

// AddContactHttpHandler handles the HTTP request for adding a contact
// @Summary Add a new contact
// @Description Add a new contact to the phone book
// @Tags contacts
// @Accept  json
// @Produce  json
// @Param contact body AddContactRequest true "Contact to add"
// @Success 201 {object} AddContactRequest
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to save contact"
// @Router /contacts [post]
func AddContactHttpHandler(repo repositories.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body into AddContactRequest
		var req AddContactRequest
		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		// Decode the request body into AddContactRequest
		if err := json.Unmarshal(body, &req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate the request
		if err := req.Validate(); err != nil {
			if _, ok := err.(domain.DomainError); ok {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, "Failed to process request", http.StatusInternalServerError)
			}
			return
		}

		// Convert request to command
		cmd, err := req.ToCommand()
		if err != nil {
			http.Error(w, "Failed to process request", http.StatusInternalServerError)
			return
		}

		// Create handler and execute
		handler := handlers.NewAddContactHandler(repo)
		if err := handler.Handle(cmd); err != nil {
			log.Println(err)
			http.Error(w, "Failed to save contact", http.StatusInternalServerError)
			return
		}

		// Respond with success
		w.WriteHeader(http.StatusCreated)
	}
}

func (r AddContactRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		var invalidFields []string
		for _, err := range err.(validator.ValidationErrors) {
			invalidFields = append(invalidFields, err.Field())
		}
		return domain.NewInvalidContactError("required fields missing: " + strings.Join(invalidFields, ", "))
	}
	return nil
}

func (r AddContactRequest) ToCommand() (commands.AddContactCommand, error) {
	var phones []model.PhoneType
	phones = append(phones, r.Phones...)

	address, err := model.NewAddress(r.Street, r.PostalCode, r.CityId, r.CountryId)
	if err != nil {
		return commands.AddContactCommand{}, err
	}

	return commands.NewAddContactCommand(r.FirstName, r.LastName, phones, address)
}
