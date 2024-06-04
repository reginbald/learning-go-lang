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
	{ID: 12, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 13, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 14, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 15, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 16, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 17, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 18, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 19, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 20, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 21, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 22, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 23, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 24, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 25, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 26, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 27, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 28, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 29, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 30, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 31, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 32, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 33, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 34, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 35, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 36, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 37, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 38, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 39, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 40, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 41, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 42, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 43, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 44, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 45, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 46, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 47, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 48, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 49, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 50, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 51, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 52, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 53, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 54, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 55, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 56, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 57, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 58, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 59, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 60, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
	{ID: 61, FirstName: "Sara", LastName: "Black", Phone: "112", Email: "sara@black.com"},
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
