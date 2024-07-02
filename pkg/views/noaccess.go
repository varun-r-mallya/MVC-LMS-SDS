package views

import (
	"net/http"
	"html/template"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
)

func NoAccess(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./tmpl/login/noaccess.html")
	if err != nil {
		neem.Spotlight(err, "could not parse noaccess template")
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		neem.Spotlight(err, "could not serve noaccess")
	}
}

func NoBook(w http.ResponseWriter, r *http.Request, title string) {
	tmpl, err := template.ParseFiles("./tmpl/search/booknotfound.html")
	if err != nil {
		neem.Spotlight(err, "could not parse nobook template")
	}
	data := struct {
		Title string
	}{
		Title: title,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		neem.Spotlight(err, "could not serve nobook")
	}
}