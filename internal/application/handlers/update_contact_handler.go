package handlers

import (
	"github.com/niv-e/phonebook-api/internal/application/commands"
	"github.com/niv-e/phonebook-api/internal/domain/repositories"
)

type UpdateContactHandler struct {
	repo repositories.ContactRepository
}

func NewUpdateContactHandler(repo repositories.ContactRepository) *UpdateContactHandler {
	return &UpdateContactHandler{repo: repo}
}

func (h *UpdateContactHandler) Handle(cmd commands.UpdateContactCommand) error {
	contact, err := h.repo.FindByID(cmd.ID)
	if err != nil {
		return err
	}

	contact.FirstName = cmd.FirstName
	contact.LastName = cmd.LastName
	contact.Phones = cmd.Phones
	contact.Address.Street = cmd.Street
	contact.Address.CityId = cmd.CityId
	contact.Address.PostalCode = cmd.PostalCode
	contact.Address.CountryId = cmd.CountryId

	return h.repo.Update(contact)
}
