package handlers

import (
	"github.com/niv-e/phonebook-api/internal/application/model"
	"github.com/niv-e/phonebook-api/internal/application/queries"
	"github.com/niv-e/phonebook-api/internal/domain/repositories"
)

type SearchContactHandler struct {
	repo repositories.ContactRepository
}

func NewSearchContactHandler(repo repositories.ContactRepository) *SearchContactHandler {
	return &SearchContactHandler{repo: repo}
}

func (h *SearchContactHandler) Handle(query queries.SearchContactQuery) ([]model.ContactType, error) {
	return h.repo.Search(query.FirstName, query.LastName, query.FullName, query.Phone)
}
