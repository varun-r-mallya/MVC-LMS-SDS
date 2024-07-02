package controllers

import (
	// "fmt"
	"net/http"
	"encoding/json"
	"io"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/models"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/passwords"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/controllers/jsonwebtoken"
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

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Admin login API accessed")
	var user types.UserLogin
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.Unmarshal(body, &user)
	if !user.IsAdmin {
		toSend := types.Message{Message: "Not an admin"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusUnauthorized)
		return
	}
	admin, err := models.GetUser(user.UserName)
	if err != nil {
		toSend := types.Message{Message: "User does not exist"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusUnauthorized)
		return
	}
	if !admin.IsAdmin {
		toSend := types.Message{Message: "Not an admin"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusUnauthorized)
		return
	}
	hashed_password := passwords.ComparePasswords(user.Password, admin.HashedPassword, admin.Salt)
	if !hashed_password {
		toSend := types.Message{Message: "Incorrect password"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusUnauthorized)
		return
	}
	toSendUser := types.CookieUser{UserName: user.UserName, IsAdmin: true}
	w = jsonwebtoken.SetCookieHandler(w, toSendUser, "/")
	http.Redirect(w, r, "/admin/dashboard", http.StatusFound)

}

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Admin dashboard accessed")
	Librarydata, err := models.GetLibraryData()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	AdminData, err := models.GetCheckRequests()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	pageDataAdmin := types.PageDataAdmin{
		LibraryData: Librarydata,
		AdminData: AdminData,
	}
	views.AdminDashboard(w, r, pageDataAdmin)
}