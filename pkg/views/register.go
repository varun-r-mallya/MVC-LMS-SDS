package views

import (
	"net/http"
	"html/template"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
)

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./tmpl/login/register.html")
	if err != nil {
		neem.Spotlight(err, "could not parse homepage template")
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		neem.Spotlight(err, "could not serve homepage")
	}
}