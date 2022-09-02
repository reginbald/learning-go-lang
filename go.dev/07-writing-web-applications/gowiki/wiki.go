package main

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var pageNameSearch = regexp.MustCompile("\\[([a-zA-Z0-9]+)\\]")

const dir = "data/"

func init() {
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	http.HandleFunc("/", frontPageHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":4567", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2]) // The tile is the second subexpression
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {

	html := new(bytes.Buffer)
	err := templates.ExecuteTemplate(html, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var response string
	switch {
	case strings.Contains(tmpl, "view"):
		response = pageNameSearch.ReplaceAllString(html.String(), "<a href=\"/view/${1}\">${1}</a>")
	default:
		response = html.String()
	}
	_, err = io.WriteString(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *Page) save() error {

	filename := dir + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
