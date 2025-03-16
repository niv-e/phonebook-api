package main

import (
	"log"
	"net/http"

	_ "github.com/niv-e/phonebook-api/docs" // This is required to load the generated docs
	api "github.com/niv-e/phonebook-api/internal/delivery/http/endpoint"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Phone Book API
// @version 1.0
// @description A simple RESTful API for managing a phone book.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	http.HandleFunc("/contacts", api.AddContactHttpHandler)

	// Serve the Swagger UI
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
