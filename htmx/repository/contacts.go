package repository

import "strings"

var contacts = []Contact{
	{ID: 1, FirstName: "Joe", LastName: "Smith", Phone: "911", Email: "joe@smith.com"},
	{ID: 2, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
}

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

func GetContacts() []Contact {
	return contacts
}

func AddContact(c Contact) {
	c.ID = len(contacts) + 1
	contacts = append(contacts, c)
}

func GetContact(query string) []Contact {
	out := make([]Contact, 0)
	for _, c := range contacts {
		if strings.EqualFold(c.FirstName, query) {
			out = append(out, c)
		}
	}

	return out
}
