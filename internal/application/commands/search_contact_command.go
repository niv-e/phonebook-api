package commands

type SearchContactCommand struct {
	FirstName string
	LastName  string
	FullName  string
	Phone     string
}

func NewSearchContactCommand(firstName, lastName, fullName, phone string) SearchContactCommand {
	return SearchContactCommand{
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullName,
		Phone:     phone,
	}
}
