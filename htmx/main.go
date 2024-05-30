package main

import (
	"net/http"

	"github.com/reginbald/learning-go-lang/htmx/views/contacts"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("./assets/"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/contacts", http.StatusSeeOther) })

	r.Route("/contacts", contacts.ContactsRouter)

	http.ListenAndServe(":3000", r)
}
