package api

import (
	"net/http"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/controllers"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/controllers/jsonwebtoken"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
)

func Server(){
	http.HandleFunc("/", controllers.Homepage)
	http.HandleFunc("/admin", controllers.Admin)
	http.HandleFunc("/client", controllers.Client)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/noaccess", controllers.NoAccess)

	http.HandleFunc("/api/register", controllers.RegisterUser)
	http.HandleFunc("/admin/api/login", controllers.AdminLogin)
	http.HandleFunc("/client/api/login", controllers.ClientLogin)

	http.HandleFunc("/admin/dashboard", jsonwebtoken.Middleware("/admin/dashboard", controllers.AdminDashboard))
	http.HandleFunc("/admin/api/addbooks", jsonwebtoken.Middleware("/admin/api/addbooks", controllers.AddBooks))

	http.HandleFunc("/client/dashboard", jsonwebtoken.Middleware("/client/dashboard", controllers.ClientDashboard))
	http.HandleFunc("/client/api/requestadmin", jsonwebtoken.Middleware("/client/api/requestadmin", controllers.RequestAdmin))

	neem.Log("Find page at http://localhost:8080/ or http://xeonlib.org")
	neem.Critial(http.ListenAndServe(":8080", nil), "Error starting server")
}