package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/controllers/jsonwebtoken"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/models"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
)


func RequestAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
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
	err1 := models.FlipAdmin(user)
	if err1 != nil {
		message := types.Message{Message: "Could Not send Request"}
		b, err := json.Marshal(message)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusNotModified)
		return
	}
	message := types.Message{Message: "Request Sent"}
	b, err := json.Marshal(message)
	if err != nil {
		neem.Spotlight(err, "could not marshal message")
	}
	w.Write(b)


	
	
	
}

