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
	http.HandleFunc("/admin/viewbook", jsonwebtoken.Middleware("/admin/viewbook", controllers.AdminViewBook))
	http.HandleFunc("/admin/api/updatebooks", jsonwebtoken.Middleware("/admin/api/updatebooks", controllers.UpdateBooks))
	http.HandleFunc("/admin/api/deletebooks", jsonwebtoken.Middleware("/admin/api/deletebooks", controllers.DeleteBooks))
	http.HandleFunc("/admin/api/handlecheckouts", jsonwebtoken.Middleware("/admin/api/handlecheckouts", controllers.AcceptCheckOut))
	http.HandleFunc("/admin/api/handlecheckins", jsonwebtoken.Middleware("/admin/api/handlecheckins", controllers.AcceptCheckIn))
	http.HandleFunc("/admin/api/manageadmins", jsonwebtoken.Middleware("/admin/api/manageadmins", controllers.AcceptAdmins))

	http.HandleFunc("/client/dashboard", jsonwebtoken.Middleware("/client/dashboard", controllers.ClientDashboard))
	http.HandleFunc("/client/viewbook", jsonwebtoken.Middleware("/client/viewbook", controllers.ClientViewBook))
	http.HandleFunc("/client/api/checkout", jsonwebtoken.Middleware("/client/api/checkout", controllers.HandleCheckOut))
	http.HandleFunc("/client/api/checkin", jsonwebtoken.Middleware("/client/api/checkin", controllers.HandleCheckIn))
	http.HandleFunc("/client/api/requestadmin", jsonwebtoken.Middleware("/client/api/requestadmin", controllers.RequestAdmin))

	neem.Log("Find page at http://localhost:8080/ or http://xeonlib.org")
	neem.Critial(http.ListenAndServe(":8080", nil), "Error starting server")
}