package handlers

import (
    "github.com/niv-e/phonebook-api/internal/application/commands"
    "github.com/niv-e/phonebook-api/internal/domain/repositories"
)

type DeleteContactHandler struct {
    repo repositories.ContactRepository
}

func NewDeleteContactHandler(repo repositories.ContactRepository) *DeleteContactHandler {
    return &DeleteContactHandler{repo: repo}
}

func (h *DeleteContactHandler) Handle(cmd commands.DeleteContactCommand) error {
    return h.repo.Delete(cmd.ID)
}