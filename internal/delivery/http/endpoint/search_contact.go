package http

import (
	"encoding/json"
	"net/http"

	"github.com/niv-e/phonebook-api/internal/application/handlers"
	"github.com/niv-e/phonebook-api/internal/application/queries"
	"github.com/niv-e/phonebook-api/internal/domain/repositories"
)

// SearchContactHttpHandler handles the HTTP request for searching contacts
// @Summary Search contacts
// @Description Search contacts by first name, last name, full name, or phone number
// @Tags contacts
// @Accept  json
// @Produce  json
// @Param first_name query string false "First Name"
// @Param last_name query string false "Last Name"
// @Param full_name query string false "Full Name"
// @Param phone query string false "Phone Number"
// @Success 200 {array} model.ContactType
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to search contacts"
// @Router /contacts/search [get]
func SearchContactHttpHandler(repo repositories.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")
		fullName := r.URL.Query().Get("full_name")
		phone := r.URL.Query().Get("phone")

		query := queries.NewSearchContactQuery(firstName, lastName, fullName, phone)
		handler := handlers.NewSearchContactHandler(repo)
		contacts, err := handler.Handle(query)
		if err != nil {
			http.Error(w, "Failed to search contacts", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contacts)
	}
}
