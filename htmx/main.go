package main

import (
	"net/http"
	"os"

	"github.com/reginbald/learning-go-lang/htmx/views/contacts"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/reginbald/learning-go-lang/htmx/reload"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	if os.Getenv("ENV") == "" || os.Getenv("ENV") == "dev" {
		r.Use(reload.ReloadMiddleware)
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/contacts", http.StatusSeeOther) })

	r.Route("/contacts", func(r chi.Router) {
		r.Get("/", contacts.GetContacts)
		r.Get("/new", contacts.GetContactForm)
		r.Post("/new", contacts.PostContact)
		r.Get("/{contactID}", contacts.GetContact)
		r.Get("/{contactID}/edit", contacts.GetContactEditForm)
	})

	http.ListenAndServe(":3000", r)
}
