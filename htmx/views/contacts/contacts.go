package contacts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/reginbald/learning-go-lang/htmx/components"
	"github.com/reginbald/learning-go-lang/htmx/repository"
	"github.com/reginbald/learning-go-lang/htmx/views"
)

func ContactsRouter(r chi.Router) {
	r.Get("/", GetContacts)
	r.Get("/new", GetContactForm)
	r.Post("/new", CreateContact)
	r.Get("/{contactID}", GetContact)
	r.Get("/{contactID}/edit", GetContactEditForm)
	r.Get("/{contactID}/email", ValidateEmail)
	r.Post("/{contactID}/edit", UpdateContact)
	r.Delete("/{contactID}", DeleteContact)
}

func ValidateEmail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "contactID"))
	if err != nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}
	email := r.URL.Query().Get("email")
	err = repository.Validate(repository.Contact{ID: id, Email: email})
	invalid, errMessage := "false", ""
	if err != nil {
		invalid, errMessage = "true", err.Error()
	}

	templ.Handler(components.Input(components.InputArguments{
		Id:          "email",
		Name:        "email",
		IType:       "email",
		Placeholder: "Email",
		Value:       email,
		AriaInvalid: invalid,
		Err:         errMessage,
		Htmx: &components.HTMX{
			Get: fmt.Sprintf("/contacts/%d/email", id),
		},
	})).ServeHTTP(w, r)
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	page := 1
	if qPage := r.URL.Query().Get("page"); qPage != "" {
		iPage, err := strconv.Atoi(qPage)
		if err != nil {
			http.Redirect(w, r, "/contacts", http.StatusSeeOther)
			return
		}
		page = iPage
	}

	if query != "" {
		templ.Handler(views.Index(index(query, page, repository.SearchContacts(query, page)))).ServeHTTP(w, r)
		return
	}

	templ.Handler(views.Index(index(query, page, repository.GetContacts(page)))).ServeHTTP(w, r)
}

func GetContactForm(w http.ResponseWriter, r *http.Request) {
	errors := map[string]string{}
	contact := repository.Contact{}
	templ.Handler(views.Index(newContact(contact, errors))).ServeHTTP(w, r)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	errors := map[string]string{}
	contact := repository.Contact{}
	err := r.ParseForm()
	if err != nil {
		templ.Handler(views.Index(newContact(contact, errors))).ServeHTTP(w, r)
		return
	}

	if contact.FirstName = r.PostForm.Get("first_name"); contact.FirstName == "" {
		errors["first_name"] = "first name missing"
	}
	if contact.LastName = r.PostForm.Get("last_name"); contact.LastName == "" {
		errors["last_name"] = "last name missing"
	}
	if contact.Phone = r.PostForm.Get("phone"); contact.Phone == "" {
		errors["phone"] = "phone missing"
	}
	if contact.Email = r.PostForm.Get("email"); contact.Email == "" {
		errors["email"] = "email missing"
	}
	if len(errors) > 0 {
		templ.Handler(views.Index(newContact(contact, errors))).ServeHTTP(w, r)
		return
	}

	// Store
	repository.AddContact(contact)
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "contactID"))
	if err != nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}
	c, err := repository.GetContact(id)
	if err != nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}
	templ.Handler(views.Index(show(*c))).ServeHTTP(w, r)
}

func GetContactEditForm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "contactID"))
	if err != nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}
	c, err := repository.GetContact(id)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/contacts/%d", id), http.StatusSeeOther)
		return
	}
	errors := map[string]string{}
	templ.Handler(views.Index(edit(*c, errors))).ServeHTTP(w, r)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	errors := map[string]string{}
	contact := repository.Contact{}
	err := r.ParseForm()
	if err != nil {
		templ.Handler(views.Index(edit(contact, errors))).ServeHTTP(w, r)
		return
	}

	if contact.FirstName = r.PostForm.Get("first_name"); contact.FirstName == "" {
		errors["first_name"] = "first name missing"
	}
	if contact.LastName = r.PostForm.Get("last_name"); contact.LastName == "" {
		errors["last_name"] = "last name missing"
	}
	if contact.Phone = r.PostForm.Get("phone"); contact.Phone == "" {
		errors["phone"] = "phone missing"
	}
	if contact.Email = r.PostForm.Get("email"); contact.Email == "" {
		errors["email"] = "email missing"
	}
	if len(errors) > 0 {
		templ.Handler(views.Index(edit(contact, errors))).ServeHTTP(w, r)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "contactID"))
	if err != nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}
	_, err = repository.UpdateContact(id, contact)
	if err != nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/contacts/%d", id), http.StatusSeeOther)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "contactID"))
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/contacts/%d", id), http.StatusSeeOther)
		return
	}
	_, err = repository.DeleteContact(id)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/contacts/%d", id), http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}
