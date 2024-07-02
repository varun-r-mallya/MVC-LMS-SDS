package models

import (
	"fmt"
	
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"

	"github.com/go-sql-driver/mysql"
)

func AddBooks(book types.Book) (bool, error) {
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error %s connecting to the database")
		return false, err
	}
	insertSql := "INSERT INTO booklist (Title, Author, Genre, NumberofCopies, NumberofCopiesAvailable, NumberofCopiesBorrowed, DueTime) VALUES (?, ?, ?, ?, ?, ?, ?);"
	_, err = db.Exec(insertSql, book.Title, book.Author, book.Genre, book.Quantity, book.Quantity, 0, book.DueTime)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			neem.Log("Duplicate entry in books")
			return false, fmt.Errorf("book already exists")
		} else {
			neem.DBError("error inserting into the database", err)
			return false, fmt.Errorf("error in database")
		}
	} else {
		neem.Log("User registered successfully")
		return true, nil
	}
}