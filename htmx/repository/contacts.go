package repository

import (
	"fmt"
	"slices"
	"strings"
)

var contacts = []Contact{
	{ID: 1, FirstName: "Joe", LastName: "Smith", Phone: "911", Email: "joe@smith.com"},
	{ID: 2, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 3, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 4, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 5, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 6, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 7, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 8, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 9, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 10, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 11, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
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

func GetContacts(page int) []Contact {
	return paginate(page, contacts)
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

func SearchContacts(query string, page int) []Contact {
	out := make([]Contact, 0)
	for _, c := range contacts {
		if strings.EqualFold(c.FirstName, query) {
			out = append(out, c)
		}
	}

	return paginate(page, out)
}

func paginate(page int, list []Contact) []Contact {
	start := (page - 1) * 10
	if start > len(list) {
		start = len(list)
	}
	end := start + 10
	if end > len(list) {
		end = len(list)
	}
	return list[start:end]
}
