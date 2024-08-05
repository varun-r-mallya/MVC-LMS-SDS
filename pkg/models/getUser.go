package models

import (
	"database/sql"
	"fmt"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"

	_ "github.com/go-sql-driver/mysql"
)

func GetUser(UserName string) (types.User, error) {
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return types.User{}, err
	}
	GetUserFromUserList := "SELECT * FROM userlist WHERE username = ?"
	var user types.User
	err = db.QueryRow(GetUserFromUserList, UserName).Scan(&user.UserName, &user.HashedPassword, &user.Salt, &user.IsAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			neem.Log("User not present in database")
			return types.User{}, fmt.Errorf("user does not exist")
		} else {
			neem.DBError("error getting from database", err)
			return types.User{}, fmt.Errorf("error in database")
		}
	}

	return user, nil

}
