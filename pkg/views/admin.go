package views

import (
	"html/template"
	"net/http"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./tmpl/login/admin.html")
	if err != nil {
		neem.Spotlight(err, "could not parse Admin template")
		
	}
	neem.Spotlight(tmpl.Execute(w, nil), "could not serve Admin login")
}

func AdminDashboard(w http.ResponseWriter, r *http.Request, data types.PageDataAdmin){
	tmpl, err := template.ParseFiles("./tmpl/dashboards/admin.html")
	if err != nil {
		neem.Spotlight(err, "could not parse Admin Dashboard template")
		
	}

	neem.Spotlight(tmpl.Execute(w, data), "could not serve Admin Dashboard")
}

func AdminViewBook(w http.ResponseWriter, r *http.Request, data types.AdminBookView){
	tmpl, err := template.ParseFiles("./tmpl/search/admin.html")
	if err != nil {
		neem.Spotlight(err, "could not parse Admin View Book template")
		
	}

	if data.Book.Title == "" {
		NoBook(w, r, r.URL.Query().Get("search"))
		return
	}
	neem.Spotlight(tmpl.Execute(w, data), "could not serve Admin View Book")
}