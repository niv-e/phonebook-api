package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/niv-e/phonebook-api/internal/application/handlers"
	"github.com/niv-e/phonebook-api/internal/application/queries"
	"github.com/niv-e/phonebook-api/internal/infrastructure/persistence"
)

// GetContactsHttpHandler handles the HTTP request for getting paginated contacts
// @Summary Get paginated contacts
// @Description Get paginated contacts from the phone book
// @Tags contacts
// @Accept  json
// @Produce  json
// @Param page query int true "Page number"
// @Success 200 {array} model.ContactType
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to fetch contacts"
// @Router /contacts [get]
func GetContactsHttpHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	query, err := queries.NewGetContactsQuery(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handler := handlers.NewGetContactsHandler(persistence.GetContactRepository())
	contacts, err := handler.Handle(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}
