package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/controllers/jsonwebtoken"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/models"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
)

func cookieUserDetector(w http.ResponseWriter, r *http.Request) (types.CookieUser, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		neem.Spotlight(err, "Error in Cookie decoding")
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Redirect(w, r, "/noaccess", http.StatusFound)
			return types.CookieUser{}, errors.New("no cookie")
		default:
			neem.Spotlight(err, "Cookie error")
			http.Redirect(w, r, "/noaccess", http.StatusFound)
			return types.CookieUser{}, errors.New("cookie error")
		}
	}
	user, err2 := jsonwebtoken.ValidateToken(cookie.Value)
	if err2 != nil {
		http.Redirect(w, r, "/noaccess", http.StatusFound)
		return types.CookieUser{}, errors.New("token error")
	}
	return user, nil
}

func RequestAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	user, err := cookieUserDetector(w, r)
	if err != nil {
		return
	}
	err1 := models.FlipAdmin(user)
	if err1 != nil {
		message := types.Message{Message: "Could not send Request"}
		b, err := json.Marshal(message)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		w.Write(b)
		return
	}
	message := types.Message{Message: "Request sent to become admin"}
	b, err := json.Marshal(message)
	if err != nil {
		neem.Spotlight(err, "could not marshal message")
	}
	w.Write(b)
}

func HandleCheckOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Client checkout accessed")
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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	strbookID := struct {
		BookID string `json:"BookID"`
	}{}
	json.Unmarshal(body, &strbookID)
	bookID, err := strconv.Atoi(strbookID.BookID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	message, err := models.RequestCheckOut(user, bookID)
	if err != nil {
		toSend := types.Message{Message: message}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		w.Write(b)
		return
	}
	toSend := types.Message{Message: message}
	b, err := json.Marshal(toSend)
	if err != nil {
		neem.Spotlight(err, "could not marshal message")
	}
	w.Write(b)
}

func HandleCheckIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Client checkin accessed")
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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	strbookID := struct {
		BookID string `json:"BookID"`
	}{}
	json.Unmarshal(body, &strbookID)
	bookID, err := strconv.Atoi(strbookID.BookID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	message, err := models.RequestCheckIn(user, bookID)
	if err != nil {
		toSend := types.Message{Message: message}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		w.Write(b)
		return
	}
	toSend := types.Message{Message: message}
	b, err := json.Marshal(toSend)
	if err != nil {
		neem.Spotlight(err, "could not marshal message")
	}
	w.Write(b)
}
