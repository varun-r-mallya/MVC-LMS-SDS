package models

import (

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
)

func RegisterUser(user types.UserRegister) (error, bool) {
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error %s connecting to the database")
		return err, false
	}
	insertSql := "INSERT INTO userlist (username, hashedpassword, salt, isadmin) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(insertSql, user.UserName, user.HashedPassword, user.Salt, user.IsAdmin)
	if err != nil {
		//TODO:user registration error handling for duplicate entry
		//TODO: Write a separate unit test to fix this function
		neem.DBError("error inserting into the database", err)
		return err, false
	} else {
		neem.Log("User registered successfully")
		return nil, true
	}
}

