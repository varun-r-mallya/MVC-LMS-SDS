package controllers

import (
	"encoding/json"
	"errors"
	_ "fmt"
	"io"
	"net/http"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/controllers/jsonwebtoken"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/controllers/passwords"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/models"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/views"
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
	Librarydata, err := models.GetLibraryData()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	Books, err := models.BooksList()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	cookie, err := r.Cookie("token")
	if err != nil {
		neem.Spotlight(err, "Error in Cookie decoding")
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Redirect(w, r, "/noaccess", http.StatusFound)
			return
		default:
			neem.Spotlight(err, "Cookie error")
			http.Redirect(w, r, "/noaccess", http.StatusFound)
			return
		}
	}
	user, err2 := jsonwebtoken.ValidateToken(cookie.Value)
	if err2 != nil {
		http.Redirect(w, r, "/noaccess", http.StatusFound)
		return
	}
	transactions, err := models.ClientTransactions(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	pageDataClient := types.PageDataClient{
		LibraryData:  Librarydata,
		Books:        Books,
		Transactions: transactions,
	}
	views.ClientDashboard(w, r, pageDataClient)
}

func ClientViewBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query()
	neem.Log("Client view book accessed")
	title := query.Get("search")
	if title == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	bookdata, err := models.GetBook(title)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	cookie, err := r.Cookie("token")
	if err != nil {
		neem.Spotlight(err, "Error in Cookie decoding")
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Redirect(w, r, "/noaccess", http.StatusFound)
			return
		default:
			neem.Spotlight(err, "Cookie error")
			http.Redirect(w, r, "/noaccess", http.StatusFound)
			return
		}
	}
	user, err2 := jsonwebtoken.ValidateToken(cookie.Value)
	if err2 != nil {
		http.Redirect(w, r, "/noaccess", http.StatusFound)
		return
	}
	transactions, err := models.ClientPerBookTransactions(user, bookdata.B_Id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := types.ClientBookView{
		Book:         bookdata,
		Transactions: transactions,
	}
	views.ClientViewBook(w, r, data)
}
