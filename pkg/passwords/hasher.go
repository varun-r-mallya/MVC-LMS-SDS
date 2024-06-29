package passwords

import (
	"crypto/sha256"
	"encoding/hex"
	"os"

)

func hashPassword(password string, salt string) (string, string) {
	h := sha256.New()
	h.Write([]byte(password))
	hashed_password := hex.EncodeToString(h.Sum(nil))
	return hashed_password, salt
}


func SaltingPassword(password string, username string, user_type string) (string, string) {
	GlobalSalt := os.Getenv("GLOBALSALT")
	salt := saltgen()
	transformed_salt := (salt + GlobalSalt)
	return password+transformed_salt, salt
}

func PasswordTransform(password string, username string, user_type string) (string, string) {
	return hashPassword(SaltingPassword(password, username, user_type))
}

func ComparePasswords(password string, hashed_password string, salt string) bool {
	
	GlobalSalt := os.Getenv("GLOBALSALT")
	transformed_salt := salt + GlobalSalt
	password = password + transformed_salt
	hashed_password_new, _ := hashPassword(password, salt)
	return hashed_password_new == hashed_password
}