package controllers

import (
	"net/http"
	"io"
	"encoding/json"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/views"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/models"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/passwords"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/controllers/jsonwebtoken"


)

func Client(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Client login accessed")
	views.Client(w, r)
}

func ClientLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Client login API accessed")
	var user types.UserLogin
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.Unmarshal(body, &user)
	if user.IsAdmin {
		toSend := types.Message{Message: "User is an Admin"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusUnauthorized)
		return
	}
	client, err := models.GetUser(user.UserName)
	if err != nil {
		toSend := types.Message{Message: "User does not exist"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusUnauthorized)
		return
	}
	if client.IsAdmin {
		toSend := types.Message{Message: "User is an Admin"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusUnauthorized)
		return
	}
	hashed_password := passwords.ComparePasswords(user.Password, client.HashedPassword, client.Salt)
	if !hashed_password {
		toSend := types.Message{Message: "Incorrect password"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusUnauthorized)
		return
	}
	toSendUser := types.CookieUser{UserName: user.UserName, IsAdmin: false}
	w = jsonwebtoken.SetCookieHandler(w, toSendUser, "/")
	http.Redirect(w, r, "/client/dashboard", http.StatusFound)

}

func ClientDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Client dashboard accessed")
	views.ClientDashboard(w, r)
}