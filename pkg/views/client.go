package views

import (
	"html/template"
	"net/http"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
)

func Client(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./tmpl/login/client.html")
	if err != nil {
		neem.Spotlight(err, "could not parse client template")
		
	}
	neem.Spotlight(tmpl.Execute(w, nil), "could not serve Client login")
}

func ClientDashboard(w http.ResponseWriter, r *http.Request, data types.PageDataClient){
	tmpl, err := template.ParseFiles("./tmpl/dashboards/client.html")
	if err != nil {
		neem.Spotlight(err, "could not parse Client Dashboard template")
		
	}
	neem.Spotlight(tmpl.Execute(w, data), "could not serve Client Dashboard")
}