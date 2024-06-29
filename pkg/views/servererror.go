package views

import (
	"net/http"
	"html/template"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
)

func ServerError(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./tmpl/error/index.html")
	if err != nil {
		neem.Spotlight(err, "could not parse server error template")
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		neem.Spotlight(err, "could not serve server error")
	}
}