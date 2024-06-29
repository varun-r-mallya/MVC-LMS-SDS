package api

import (
	"net/http"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/controllers"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
)

func Server(){
	http.HandleFunc("/", controllers.Homepage)
	http.HandleFunc("/admin", controllers.Admin)
	http.HandleFunc("/client", controllers.Client)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/api/register", controllers.RegisterUser)

	neem.Log("Find page at http://localhost:3000/ or http://xeonlib.org")
	neem.Critial(http.ListenAndServe(":8080", nil), "Error starting server")
}