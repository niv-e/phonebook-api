package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// ContactEntity represents a person's basic information
type ContactEntity struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	AddressID uint           `json:"address_id"`
	Address   AddressEntity  `json:"address" gorm:"foreignKey:AddressID"`
	Phones    datatypes.JSON `json:"phones" gorm:"type:jsonb"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (ContactEntity) TableName() string {
	return "contacts"
}

// AddressEntity represents a physical location
type AddressEntity struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Street     string     `json:"street"`
	PostalCode string     `json:"postal_code"`
	CityID     uint       `json:"city_id"`
	City       CityEntity `json:"city" gorm:"foreignKey:CityID"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func (AddressEntity) TableName() string {
	return "addresses"
}

// CityEntity represents a city or town
type CityEntity struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	Name      string        `json:"name"`
	CountryID uint          `json:"country_id"`
	Country   CountryEntity `json:"country" gorm:"foreignKey:CountryID"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (CityEntity) TableName() string {
	return "cities"
}

// CountryEntity represents a country based on ISO 3166
type CountryEntity struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Alpha2Code  string    `json:"alpha2_code" gorm:"size:2;comment:ISO 3166-1 alpha-2 code"`
	Alpha3Code  string    `json:"alpha3_code" gorm:"size:3;comment:ISO 3166-1 alpha-3 code"`
	NumericCode string    `json:"numeric_code" gorm:"size:3;comment:ISO 3166-1 numeric code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (CountryEntity) TableName() string {
	return "countries"
}
