package repository

import (
	"fmt"
	"slices"
	"strings"
)

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

func Validate(x Contact) error {
	for _, c := range contacts {
		if c.ID != x.ID && c.Email == x.Email {
			return fmt.Errorf("%s is not unique", x.Email)
		}
	}
	return nil
}

func GetContacts() []Contact {
	return contacts
}

func AddContact(c Contact) {
	c.ID = len(contacts) + 1
	contacts = append(contacts, c)
}

func DeleteContact(id int) (*Contact, error) {
	i := slices.IndexFunc(contacts, func(c Contact) bool {
		return c.ID == id
	})
	if i == -1 {
		return nil, fmt.Errorf("error contact with id: %d was not found", id)
	}

	out := contacts[i]

	contacts = append(contacts[:i], contacts[i+1:]...)

	return &out, nil
}

func GetContact(id int) (*Contact, error) {
	i := slices.IndexFunc(contacts, func(c Contact) bool {
		return c.ID == id
	})
	if i == -1 {
		return nil, fmt.Errorf("error contact with id: %d was not found", id)
	}
	return &contacts[i], nil
}

func UpdateContact(id int, update Contact) (*Contact, error) {
	i := slices.IndexFunc(contacts, func(c Contact) bool {
		return c.ID == id
	})
	if i == -1 {
		return nil, fmt.Errorf("error contact with id: %d was not found", id)
	}
	update.ID = contacts[i].ID
	contacts[i] = update
	return &contacts[i], nil
}

func SearchContacts(query string) []Contact {
	out := make([]Contact, 0)
	for _, c := range contacts {
		if strings.EqualFold(c.FirstName, query) {
			out = append(out, c)
		}
	}

	return out
}
