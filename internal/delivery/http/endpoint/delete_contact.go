package http

import (
    "net/http"

    "github.com/google/uuid"
    "github.com/niv-e/phonebook-api/internal/application/commands"
    "github.com/niv-e/phonebook-api/internal/application/handlers"
    "github.com/niv-e/phonebook-api/internal/domain/repositories"
)

// DeleteContactRequest represents the request payload for deleting a contact
type DeleteContactRequest struct {
    ID uuid.UUID `json:"id" validate:"required,uuid"`
}

// DeleteContactHttpHandler handles the HTTP request for deleting a contact
// @Summary Delete a contact
// @Description Delete a contact from the phone book
// @Tags contacts
// @Accept  json
// @Produce  json
// @Param id path string true "Contact ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to delete contact"
// @Router /contacts/{id} [delete]
func DeleteContactHttpHandler(repo repositories.ContactRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        idStr := r.URL.Query().Get("id")
        id, err := uuid.Parse(idStr)
        if err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        cmd := commands.NewDeleteContactCommand(id)
        handler := handlers.NewDeleteContactHandler(repo)
        if err := handler.Handle(cmd); err != nil {
            http.Error(w, "Failed to delete contact", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}