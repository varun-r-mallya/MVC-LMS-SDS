package models

import (
	"fmt"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"

	"github.com/go-sql-driver/mysql"
)

func RegisterUser(user types.UserRegister) (bool, error) {
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error %s connecting to the database")
		return false, err
	}
	InsertIntoUserlist := "INSERT INTO userlist (username, hashedpassword, salt, isadmin) VALUES (?, ?, ?, ?)"
	CountInUserlist := "SELECT COUNT(*) FROM userlist"
	var count int
	err = db.QueryRow(CountInUserlist).Scan(&count)
	if err != nil {
		neem.DBError("error checking the database", err)
		return false, fmt.Errorf("error in database")
	}

	var isAdmin bool
	if count == 0 {
		isAdmin = true
	} else {
		isAdmin = false
	}

	_, err = db.Exec(InsertIntoUserlist, user.UserName, user.HashedPassword, user.Salt, isAdmin)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			neem.Log("Duplicate entry in registration")
			return false, fmt.Errorf("user already exists")
		} else {
			neem.DBError("error inserting into the database", err)
			return false, fmt.Errorf("error in database")
		}
	} else {
		neem.Log("User registered successfully")
		return true, nil
	}
}
