package views

import (
	"net/http"
	"html/template"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./tmpl/login/admin.html")
	if err != nil {
		neem.Spotlight(err, "could not parse Admin template")
		
	}
	neem.Spotlight(tmpl.Execute(w, nil), "could not serve Admin login")
}

func AdminDashboard(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("./tmpl/dashboards/admin.html")
	if err != nil {
		neem.Spotlight(err, "could not parse Admin Dashboard template")
		
	}
	neem.Spotlight(tmpl.Execute(w, nil), "could not serve Admin Dashboard")
}