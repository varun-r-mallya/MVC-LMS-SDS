package models

import (
	"database/sql"
	"fmt"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"

	"github.com/go-sql-driver/mysql"
)

func FlipAdmin(user types.CookieUser) error {
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return err
	}
	InsertIntoConvertQ := "INSERT INTO convertq (username) VALUES (?);"
	_, err = db.Exec(InsertIntoConvertQ, user.UserName)
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

func RequestCheckOut(user types.CookieUser, bookId int) (string, error) {
	neem.Log("Request Check Out")
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return "", err
	}
	selectNumberOfCopies := `SELECT NumberofCopiesAvailable FROM booklist WHERE B_Id = (?);`
	SelectTheNumberofMatchingTransactions := `SELECT COUNT(*) FROM transactions WHERE username = (?) AND B_Id = (?) AND ((CheckOutAccepted = 1 AND CheckInAccepted IS NULL) OR (CheckOutAccepted = 1 AND CheckInAccepted = 0) OR (CheckOutAccepted IS NULL AND CheckInAccepted IS NULL));`
	InsertIntoTransactions := `INSERT INTO transactions (username, B_Id) VALUES ((?), (?));`

	var numberofcopiesavailable int
	err = db.QueryRow(selectNumberOfCopies, bookId).Scan(&numberofcopiesavailable)
	if err != nil {
		neem.DBError("error querying the database 1", err)
		return "Error in Database", fmt.Errorf("error in database")
	}
	if numberofcopiesavailable == 0 {
		return "No Copies available", fmt.Errorf("no copies available")
	}
	var count int
	err = db.QueryRow(SelectTheNumberofMatchingTransactions, user.UserName, bookId).Scan(&count)
	if err != nil {
		neem.DBError("error querying the database 2", err)
		if err == sql.ErrNoRows {
			neem.Log("No Transaction for user present in database")
			return "No trasactions present for user", fmt.Errorf("no trasactions present for user")
		} else {
			return "Error in Database", fmt.Errorf("error in database")
		}
	}
	if count > 0 {
		return "Book already borrowed", fmt.Errorf("already checked out")
	}
	_, err = db.Exec(InsertIntoTransactions, user.UserName, bookId)
	if err != nil {
		neem.DBError("error inserting into the database 3", err)
		return "Error in Database", fmt.Errorf("error in database")
	}
	neem.Log("Book checked out successfully")
	return "Book Checked Out", nil

}

func RequestCheckIn(user types.CookieUser, bookId int) (string, error) {
	neem.Log("Request Check Out")
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return "", err
	}
	TransactionsSelector := `SELECT COUNT(*) FROM transactions WHERE username = (?) AND B_Id = (?) AND CheckOutAccepted = 1 AND CheckInAccepted IS NULL OR CheckInAccepted != 1;`
	UpdateTransactions := `UPDATE transactions SET CheckInAccepted = 0 WHERE T_Id IN (SELECT T_Id FROM (SELECT MAX(T_Id) AS T_Id FROM transactions WHERE username = ? AND B_Id = ?) AS subquery);`

	var count int
	err = db.QueryRow(TransactionsSelector, user.UserName, bookId).Scan(&count)
	if err != nil {
		neem.DBError("error querying the database 5", err)
		if err == sql.ErrNoRows {
			neem.Log("No Transaction for user present in database")
			return "No trasactions present for user", fmt.Errorf("no trasactions present for user")
		} else {
			return "Error in Database", fmt.Errorf("error in database")
		}
	}
	if count == 0 {
		return "Book Not borrowed", fmt.Errorf("book not borrowed")
	}
	_, err = db.Exec(UpdateTransactions, user.UserName, bookId)
	if err != nil {
		neem.DBError("error querying the database 6", err)
		return "Error in Database", fmt.Errorf("error in database")
	}
	neem.Log("Book checked in successfully")
	return "Book Check in request sent", nil

}
