package models

import (
	"fmt"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"

	"github.com/go-sql-driver/mysql"
)

func AddBooks(book types.Book) (bool, error) {
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return false, err
	}
	insertBookIntoBookList := "INSERT INTO booklist (Title, Author, Genre, NumberofCopies, NumberofCopiesAvailable, NumberofCopiesBorrowed, DueTime) VALUES (?, ?, ?, ?, ?, ?, ?);"
	_, err = db.Exec(insertBookIntoBookList, book.Title, book.Author, book.Genre, book.Quantity, book.Quantity, 0, book.DueTime)
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

func UpdateBooks(book types.Book) (bool, error) {
	neem.Log("Book update function called")
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return false, err
	}
	UpdateBookList := "UPDATE booklist SET Title = ?, Author = ?, Genre = ?, NumberofCopies = ?, NumberofCopiesAvailable = ?, NumberofCopiesBorrowed = ?, DueTime = ? WHERE B_Id = ?;"
	_, err = db.Exec(UpdateBookList, book.Title, book.Author, book.Genre, book.Quantity, book.Quantity, book.NumberofCopiesBorrowed, book.DueTime, book.B_Id)
	if err != nil {
		neem.DBError("error updating the database", err)
		return false, fmt.Errorf("error in database")
	} else {
		neem.Log("Book Updated successfully")
		return true, nil
	}
}

func DeleteBooks(bookID int) (bool, error) {
	neem.Log("Book delete function called")
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return false, err
	}
	CheckNumberOfCopiesBorrowed := "SELECT NumberofCopiesBorrowed FROM booklist WHERE B_Id = ?"
	var numberofcopiesborrowed int
	err = db.QueryRow(CheckNumberOfCopiesBorrowed, bookID).Scan(&numberofcopiesborrowed)
	if err != nil {
		neem.DBError("error checking the database", err)
		return false, fmt.Errorf("error in database")
	}
	if numberofcopiesborrowed > 0 {
		neem.Log("Book has been borrowed")
		return false, fmt.Errorf("book has been borrowed")
	}
	DeleteTransaction := "DELETE FROM transactions WHERE B_Id = ?"
	_, err = db.Exec(DeleteTransaction, bookID)
	if err != nil {
		neem.DBError("error checking the database", err)
		return false, fmt.Errorf("error in database")
	}
	DeleteBookFinal := "DELETE FROM booklist WHERE B_Id = ?"
	_, err = db.Exec(DeleteBookFinal, bookID)
	if err != nil {
		neem.DBError("error deleting from the database", err)
		return false, fmt.Errorf("error in database")
	} else {
		neem.Log("Book deleted successfully")
		return true, nil
	}
}

func AcceptCheckOut(checkout types.CheckOut) (string, error) {
	neem.Log("Accept checkout function called")
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return "Error connecting to the database", err
	}
	if !checkout.Accepted {
		RemoveCheckout := "UPDATE transactions SET DateBorrowed = CURDATE(), CheckOutAccepted = 0 WHERE T_Id = ?;"
		_, err := db.Query(RemoveCheckout, checkout.T_Id)
		if err != nil {
			neem.DBError("error updating the database", err)
			return "Error updating the database", err
		}
		return "Checkout rejected", nil
	} else {
		AcceptCheckoutQuery := "UPDATE transactions SET DateBorrowed = CURDATE(), CheckOutAccepted = 1 WHERE T_Id = ?;"
		EnsureParityQuery := "UPDATE booklist SET NumberofCopiesBorrowed = NumberofCopiesBorrowed + 1, NumberofCopiesAvailable = NumberofCopiesAvailable - 1 WHERE B_Id = (SELECT B_Id FROM transactions WHERE T_Id = (?));"
		_, err := db.Query(AcceptCheckoutQuery, checkout.T_Id)
		if err != nil {
			neem.DBError("error updating the database", err)
			return "Error updating the database", err
		}
		_, err = db.Query(EnsureParityQuery, checkout.T_Id)
		if err != nil {
			neem.DBError("error updating the database", err)
			return "Error updating the database", err
		}
		return "Checkout accepted", nil
	}
}

func AcceptCheckIn(checkout types.CheckIn) (string, error) {
	neem.Log("Accept checkin function called")
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return "Error connecting to the database", err
	}
	if !checkout.Accepted {
		RemoveCheckoutQuery := "UPDATE transactions SET CheckInAccepted = NULL WHERE T_Id = ?;"
		_, err := db.Query(RemoveCheckoutQuery, checkout.T_Id)
		if err != nil {
			neem.DBError("error updating the database", err)
			return "Error updating the database", err
		}
		return "Checkin rejected", nil
	} else {
		AcceptCheckOutQuery := "UPDATE transactions SET DateReturned = CURDATE(), CheckInAccepted = 1 WHERE T_Id = ?;"
		EnsureParityQuery := "UPDATE booklist SET NumberofCopiesAvailable = NumberofCopiesAvailable + 1, NumberofCopiesBorrowed = NumberofCopiesBorrowed - 1 WHERE B_Id = (SELECT B_Id FROM transactions WHERE T_Id = ?);"
		//updateSql3 := "UPDATE TRANSACTIONS SET OverDueFine = ((DATEDIFF(CURDATE(), DateBorrowed) -  (SELECT NumberofDays FROM BOOKLIST WHERE B_Id = (SELECT B_Id FROM TRANSACTIONS WHERE T_Id = ${T_Id}))) * ${process.env.FINEPERDAY} WHERE T_Id = ${T_Id} AND DATEDIFF(CURDATE(), DateBorrowed) > (SELECT NumberofDays FROM BOOKLIST WHERE B_Id =  (SELECT B_Id FROM TRANSACTIONS WHERE T_Id = ${T_Id}));"
		_, err := db.Query(AcceptCheckOutQuery, checkout.T_Id)
		if err != nil {
			neem.DBError("error updating the database", err)
			return "Error updating the database", err
		}
		_, err = db.Query(EnsureParityQuery, checkout.T_Id)
		if err != nil {
			neem.DBError("error updating the database", err)
			return "Error updating the database", err
		}
		return "Checkin accepted", nil
	}
}

func AcceptAdmins(accept types.AcceptAdmins) (string, error) {
	neem.Log("Accept admins function called")
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return "Error connecting to the database", err
	}
	if !accept.Accepted {
		ConvertFromAdminConverterList := "DELETE FROM convertq WHERE username = ?"
		_, err := db.Query(ConvertFromAdminConverterList, accept.UserName)
		if err != nil {
			neem.DBError("error updating the database", err)
			return "Error updating the database", err
		}
		return "Admin not accepted", nil
	} else {
		UpdateUserList := "UPDATE userlist SET isadmin = 1 WHERE username = ?;"
		_, err := db.Query(UpdateUserList, accept.UserName)
		if err != nil {
			neem.DBError("error updating the database", err)
			return "Error updating the database", err
		}
		DeleteFromConverterQueue := "DELETE FROM convertq WHERE username = ?"
		_, err1 := db.Query(DeleteFromConverterQueue, accept.UserName)
		if err1 != nil {
			neem.DBError("error updating the database", err)
			return "Error updating the database", err
		}
		return "Admin added", nil
	}
}
