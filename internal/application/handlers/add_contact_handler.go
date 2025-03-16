package handlers

import (
	"github.com/niv-e/phonebook-api/internal/application/commands"
	"github.com/niv-e/phonebook-api/internal/domain/repositories"
)

type AddContactHandler struct {
	repo repositories.ContactRepository
}

func NewAddContactHandler(repo repositories.ContactRepository) *AddContactHandler {
	return &AddContactHandler{repo: repo}
}

func (h *AddContactHandler) Handle(cmd commands.AddContactCommand) error {
	contact, err := cmd.ToContactType()
	if err != nil {
		return err
	}
	return h.repo.Save(contact)
}
