package queries

type SearchContactQuery struct {
	FirstName string
	LastName  string
	FullName  string
	Phone     string
}

func NewSearchContactQuery(firstName, lastName, fullName, phone string) SearchContactQuery {
	return SearchContactQuery{
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullName,
		Phone:     phone,
	}
}
