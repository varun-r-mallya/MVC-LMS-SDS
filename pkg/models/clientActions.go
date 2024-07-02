package models

import (
	"fmt"
	
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"

	"github.com/go-sql-driver/mysql"
)

func FlipAdmin(user types.CookieUser) (error) {
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error %s connecting to the database")
		return err
	}
	insertSql := "INSERT INTO convertq (username) VALUES (?);"
	_, err = db.Exec(insertSql, user.UserName)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			neem.Log("Duplicate entry for user")
			return fmt.Errorf("request already sent")
		} else {
			neem.DBError("error inserting into the database", err)
			return fmt.Errorf("error in database")
		}
	} else {
		neem.Log("User registered successfully")
		return nil
	}
}