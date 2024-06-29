package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/models"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/views"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Registration page accessed")
	views.Register(w, r)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("User registration accessed")

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	confirm_password := r.PostFormValue("confirm_password")

	if username == "" || password == "" || confirm_password == "" {
		neem.Log("Empty fields error")
		toSend := types.Message{Message: "Empty fields"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return
	}

	if password != confirm_password {
		neem.Log("Passwords do not match error")
		toSend := types.Message{Message: "Passwords do not match"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return
	}

	register := types.UserRegister{
		UserName:       username,
		HashedPassword: password,
		Salt:           password,
		IsAdmin:        true,
	}

	err, success := models.RegisterUser(register)
	if err != nil {
		neem.Spotlight(err, "could not register user")
		toSend := types.Message{Message: "Could not register user"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusInternalServerError)
		return
	}

	if success {
		neem.Log("User registered successfully")
		toSend := types.Message{Message: "User registered successfully"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusOK)
		return
	}





}