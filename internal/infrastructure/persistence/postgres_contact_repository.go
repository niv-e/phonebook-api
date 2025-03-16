package persistence

import (
	"log"
	"os"
	"sync"

	"github.com/niv-e/phonebook-api/internal/domain"
	"github.com/niv-e/phonebook-api/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	contactRepository *PostgresContactRepository
	once              sync.Once
)

type PostgresContactRepository struct {
	db *gorm.DB
}

func GetContactRepository() *PostgresContactRepository {
	once.Do(func() {
		dsn := os.Getenv("DB_DSN")
		if dsn == "" {
			log.Fatal("DB_DSN environment variable is not set")
		}

		repo, err := NewPostgresContactRepository(dsn)
		if err != nil {
			log.Fatalf("Failed to create contact repository: %v", err)
		}
		contactRepository = repo
	})
	return contactRepository
}
func NewPostgresContactRepository(dsn string) (*PostgresContactRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresContactRepository{db: db}, nil
}
func (r *PostgresContactRepository) Save(contact entity.ContactEntity) error {
	tx := r.db.Begin()

	// Handle Country: Avoid duplicates
	var country entity.CountryEntity
	if err := tx.Where("name = ?", contact.Address.City.Country.Name).FirstOrCreate(&country, contact.Address.City.Country).Error; err != nil {
		tx.Rollback()
		return domain.NewInvalidContactError("failed to save or find country: " + err.Error())
	}
	contact.Address.City.CountryID = country.ID

	// Handle City: Avoid duplicates
	var city entity.CityEntity
	if err := tx.Where("name = ? AND country_id = ?", contact.Address.City.Name, country.ID).FirstOrCreate(&city, contact.Address.City).Error; err != nil {
		tx.Rollback()
		return domain.NewInvalidContactError("failed to save or find city: " + err.Error())
	}
	contact.Address.CityID = city.ID

	// Save Address
	if err := tx.Create(&contact.Address).Error; err != nil {
		tx.Rollback()
		return domain.NewInvalidContactError("failed to save address: " + err.Error())
	}

	// Save Contact (ID is generated here)
	if err := tx.Create(&contact).Error; err != nil {
		tx.Rollback()
		return domain.NewInvalidContactError("failed to save contact: " + err.Error())
	}

	// Assign ContactID to Phones
	for i := range contact.Phones {
		contact.Phones[i].ContactID = *contact.ID // Assuming ID is *string (UUID)
	}

	// Batch insert Phones
	if len(contact.Phones) > 0 {
		if err := tx.CreateInBatches(contact.Phones, 100).Error; err != nil { // Batch size of 100
			tx.Rollback()
			return domain.NewInvalidContactError("failed to save phones: " + err.Error())
		}
	}

	return tx.Commit().Error
}
