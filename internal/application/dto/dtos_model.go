package dto

// ContactDTO represents the simplified view of a contact for queries
type ContactDTO struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string // Convenience field for UI display
	Address   AddressDTO
	Phones    []PhoneDTO
}

// AddressDTO represents the simplified view of an address for queries
type AddressDTO struct {
	Street     string
	PostalCode string
	City       string
	Country    string
	CountryCode string // Alpha-2 code for UI display
}

// PhoneDTO represents the simplified view of a phone for queries
type PhoneDTO struct {
	Number string
	Type   string
}

// ContactSummaryDTO represents a minimal view of a contact for list displays
type ContactSummaryDTO struct {
	ID       string
	FullName string
	Country  string
	Phone    string // Primary phone number
}

// PagedResultDTO provides pagination information with results
type PagedResultDTO struct {
	Items      interface{}
	TotalCount int
	Page       int
	Limit      int
	TotalPages int
}