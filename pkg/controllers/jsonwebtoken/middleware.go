package jsonwebtoken

import (
	"errors"
	"net/http"
	"strings"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
)

func SetCookieHandler(w http.ResponseWriter, user types.CookieUser, path string) http.ResponseWriter {
	cookie := http.Cookie{
		Name:     "token",
		Value:    CreateToken(user),
		Path:     path,
		HttpOnly: false,
		Secure:   false, //TODO:change on prod
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	return w
}

func GetCookieHandler(r *http.Request) (bool, error) {
	neem.Log("Get cookie handler called")
	cookie, err := r.Cookie("token")
	if err != nil {
		neem.Spotlight(err, "Error in Cookie decoding")
		switch {
		case errors.Is(err, http.ErrNoCookie):
			return false, errors.New("no cookie found")
		default:
			neem.Spotlight(err, "Cookie error")
			return false, errors.New("internal server error")
		}
	}
	user, err2 := ValidateToken(cookie.Value)
	if err2 != nil {
		return false, errors.New("invalid token")
	}
	if user.IsAdmin && strings.HasPrefix(r.URL.Path, "/admin") {
		return true, nil
	} else if !user.IsAdmin && strings.HasPrefix(r.URL.Path, "/client") {
		return true, nil
	} else {
		return false, errors.New("no access")
	}
}

func Middleware(path string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validity, err := GetCookieHandler(r)
		if err != nil {
			http.Redirect(w, r, "/noaccess", http.StatusFound)
			return
		}
		if !validity {
			http.Redirect(w, r, "/noaccess", http.StatusFound)
			return
		} else {
			next(w, r)
		}
	}
}
