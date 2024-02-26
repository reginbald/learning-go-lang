package contacts

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/reginbald/learning-go-lang/htmx/views"
)

var contacts = []Contact{
	{id: 1, first: "Joe", last: "Smith", phone: "911", email: "joe@smith.com"},
	{id: 2, first: "Sara", last: "Black", phone: "112", email: "sara@black.com"},
}

func Contacts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	if query != "" {
		templ.Handler(views.Index(index(query, search(contacts, query)))).ServeHTTP(w, r)
		return
	}

	templ.Handler(views.Index(index(query, contacts))).ServeHTTP(w, r)
}

type Contact struct {
	id    int
	first string
	last  string
	phone string
	email string
}

func search(contacts []Contact, query string) []Contact {
	out := make([]Contact, 0)
	for _, c := range contacts {
		if strings.EqualFold(c.first, query) {
			out = append(out, c)
		}
	}

	return out
}
