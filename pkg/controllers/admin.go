package controllers

import (
	// "fmt"
	"net/http"
	"encoding/json"
	"io"
	"errors"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/models"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/controllers/passwords"
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

func AdminViewBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query()
	neem.Log("Admin view book accessed")
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
	data := types.AdminBookView{
		Book: bookdata,
		Transactions: transactions,
	}
	
	views.AdminViewBook(w, r, data)
}