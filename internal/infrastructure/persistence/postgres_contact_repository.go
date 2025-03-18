package persistence

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/niv-e/phonebook-api/internal/application/model"
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

func (r *PostgresContactRepository) Save(contact model.ContactType) error {
	phonesJSON, err := json.Marshal(contact.Phones)
	if err != nil {
		return err
	}

	contactEntity := entity.ContactEntity{
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
		Phones: phonesJSON,
	}
	log.Printf("Saving contact: %v", contactEntity)
	log.Printf("object address: %v", &contactEntity)
	return r.db.Save(&contactEntity).Error
}

func (r *PostgresContactRepository) FindPaginated(page, pageSize int) ([]model.ContactType, error) {
	var contacts []entity.ContactEntity
	offset := (page - 1) * pageSize
	if err := r.db.Preload("Address.City.Country").Limit(pageSize).Offset(offset).Find(&contacts).Error; err != nil {
		return nil, err
	}

	var contactDTOs []model.ContactType
	for _, contact := range contacts {
		var phones []model.PhoneType
		if err := json.Unmarshal([]byte(contact.Phones), &phones); err != nil {
			return nil, err
		}

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
			Phones: phones,
		})
	}

	return contactDTOs, nil
}

func (r *PostgresContactRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.ContactEntity{ID: id}).Error
}

func (r *PostgresContactRepository) FindByID(id uuid.UUID) (model.ContactType, error) {
	var contact entity.ContactEntity
	if err := r.db.Preload("Address.City.Country").First(&contact, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.ContactType{}, domain.NewInvalidContactError("contact not found")
		}
		return model.ContactType{}, err
	}

	var phones []model.PhoneType
	if err := json.Unmarshal(contact.Phones, &phones); err != nil {
		return model.ContactType{}, err
	}

	return model.ContactType{
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
		Phones: phones,
	}, nil
}

func (r *PostgresContactRepository) Update(contact model.ContactType) error {
	phonesJSON, err := json.Marshal(contact.Phones)
	if err != nil {
		return err
	}

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
		Phones: phonesJSON,
	}
	return r.db.Save(&contactEntity).Error
}

func (r *PostgresContactRepository) Search(firstName, lastName, fullName, phone string) ([]model.ContactType, error) {
	var contacts []entity.ContactEntity
	query := r.db.Preload("Address.City.Country")

	if firstName != "" {
		query = query.Where("first_name ILIKE ?", "%"+firstName+"%")
	}
	if lastName != "" {
		query = query.Where("last_name ILIKE ?", "%"+lastName+"%")
	}
	if fullName != "" {
		query = query.Where("CONCAT(first_name, ' ', last_name) ILIKE ?", "%"+fullName+"%")
	}
	if phone != "" {
        query = query.Where("EXISTS (SELECT 1 FROM jsonb_array_elements(phones) AS p WHERE p->>'number' ILIKE ?)", "%"+phone+"%")
    }

	if err := query.Find(&contacts).Error; err != nil {
		return nil, err
	}

	var contactDTOs []model.ContactType
	for _, contact := range contacts {
		var phones []model.PhoneType
		if err := json.Unmarshal(contact.Phones, &phones); err != nil {
			return nil, err
		}

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
			Phones: phones,
		})
	}

	return contactDTOs, nil
}
