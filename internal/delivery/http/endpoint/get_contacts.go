package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/niv-e/phonebook-api/internal/application/handlers"
	"github.com/niv-e/phonebook-api/internal/application/queries"
	"github.com/niv-e/phonebook-api/internal/domain/repositories"
)

// var (
// 	tracer = otel.Tracer("phonebook-api")
// 	meter  = otel.Meter("phonebook-api")
// )

// var requestCounter metric.Int64Counter

// func init() {
// 	var err error
// 	requestCounter, err = meter.Int64Counter("http.requests",
// 		metric.WithDescription("The number of HTTP requests"),
// 		metric.WithUnit("{request}"))
// 	if err != nil {
// 		panic(err)
// 	}
// }

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
func GetContactsHttpHandler(repo repositories.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx, span := tracer.Start(r.Context(), "GetContactsHttpHandler")
		// defer span.End()

		// requestCounter.Add(ctx, 1, metric.WithAttributes(attribute.String("http.route", "/contacts")))

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

		handler := handlers.NewGetContactsHandler(repo)
		contacts, err := handler.Handle(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contacts)
	}
}
