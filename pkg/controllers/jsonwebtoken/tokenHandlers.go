package jsonwebtoken

// THIS PACKAGE RELIES ON ENVIROMENT VARIABLES BEING PRELOADED INTO THE ENVIRONMENT. MAKE SURE THEY ARE PRELOADED
// BEFORE RUNNING THIS FUNCTION
import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
)

func CreateToken(user types.CookieUser) string {

	secret := os.Getenv("JWTSECRET")
	expiration_set, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		neem.Spotlight(err, "Error converting string to int")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	expirationTime := time.Now().Add(time.Minute * time.Duration(expiration_set)).Unix()

	claims["user"] = user.UserName
	claims["isadmin"] = user.IsAdmin
	claims["exp"] = expirationTime

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		neem.Spotlight(err, "Error Signing token")
	}

	return tokenString
}

func ValidateToken(tokenString string) (types.CookieUser, error) {
	neem.Log("Validating token")
	secret := os.Getenv("JWTSECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		neem.Log("Error parsing token")
		return types.CookieUser{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userData := types.CookieUser{
			UserName: claims["user"].(string),
			IsAdmin:  claims["isadmin"].(bool),
		}
		neem.Log(fmt.Sprintf("%s %v logged in", userData.UserName, userData.IsAdmin))
		return userData, nil
	} else {
		neem.Log("Invalid token")
		return types.CookieUser{}, fmt.Errorf("invalid token")
	}
}
