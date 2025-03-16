package persistence

import (
	"log"
	"os"
	"sync"

	"github.com/niv-e/phonebook-api/internal/application/model"
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

func (r *PostgresContactRepository) Save(contact model.ContactType) error {
	contactEntity := entity.ContactEntity{
		ID:        *contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Address: entity.AddressEntity{
			Street:     contact.Address.Street,
			PostalCode: contact.Address.PostalCode,
			City: entity.CityEntity{
				ID: contact.Address.CityId,
				Country: entity.CountryEntity{
					ID: contact.Address.CountryId,
				},
			},
		},
		Phones: convertToPhoneEntities(contact.Phones),
	}
	return r.db.Save(&contactEntity).Error
}

func (r *PostgresContactRepository) FindPaginated(page, pageSize int) ([]model.ContactType, error) {
	var contacts []entity.ContactEntity
	offset := (page - 1) * pageSize
	if err := r.db.Preload("Address.City.Country").Preload("Phones").Limit(pageSize).Offset(offset).Find(&contacts).Error; err != nil {
		return nil, err
	}

	var contactDTOs []model.ContactType
	for _, contact := range contacts {
		contactDTOs = append(contactDTOs, model.ContactType{
			ID:        &contact.ID,
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Address: model.AddressType{
				Street:      contact.Address.Street,
				PostalCode:  contact.Address.PostalCode,
				CityId:      contact.Address.City.ID,
				CityName:    contact.Address.City.Name,
				CountryId:   contact.Address.City.Country.ID,
				CountryName: contact.Address.City.Country.Name,
			},
			Phones: convertToPhoneDTOs(contact.Phones),
		})
	}

	return contactDTOs, nil
}

func convertToPhoneEntities(phones []model.PhoneType) []entity.PhoneEntity {
	var phoneEntities []entity.PhoneEntity
	for _, phone := range phones {
		phoneEntities = append(phoneEntities, entity.PhoneEntity{
			Number: phone.Number,
			Type:   phone.Type,
		})
	}
	return phoneEntities
}

func convertToPhoneDTOs(phones []entity.PhoneEntity) []model.PhoneType {
	var phoneDTOs []model.PhoneType
	for _, phone := range phones {
		phoneDTOs = append(phoneDTOs, model.PhoneType{
			Number: phone.Number,
			Type:   phone.Type,
		})
	}
	return phoneDTOs
}
