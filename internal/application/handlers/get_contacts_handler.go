package handlers

import (
	"github.com/niv-e/phonebook-api/internal/application/model"
	"github.com/niv-e/phonebook-api/internal/application/queries"
	"github.com/niv-e/phonebook-api/internal/domain/repositories"
)

type GetContactsHandler struct {
	repo repositories.ContactRepository
}

func NewGetContactsHandler(repo repositories.ContactRepository) *GetContactsHandler {
	return &GetContactsHandler{repo: repo}
}

func (h *GetContactsHandler) Handle(query queries.GetContactsQuery) ([]model.ContactType, error) {
	return h.repo.FindPaginated(query.Page, 10)
}
