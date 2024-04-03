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
	// if os.Getenv("ENV") == "" || os.Getenv("ENV") == "dev" {
	// 	r.Use(reload.ReloadMiddleware)
	// }

	r.Get("/", func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/contacts", http.StatusSeeOther) })

	r.Route("/contacts", contacts.ContactsRouter)

	http.ListenAndServe(":3000", r)
}
