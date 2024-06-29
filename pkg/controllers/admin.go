package controllers

import (
	"net/http"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/views"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Admin login accessed")
	views.Admin(w, r)
}