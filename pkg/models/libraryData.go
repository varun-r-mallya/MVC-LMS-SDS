package models

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"

	_"github.com/go-sql-driver/mysql"
)

func GetLibraryData() (types.LibraryData, error) {
	neem.Log("Getting library data")
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return types.LibraryData{}, err
	}
	var getSql string
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM booklist").Scan(&count)
	if err != nil {
		neem.DBError("error getting column count", err)
		return types.LibraryData{}, fmt.Errorf("error in database")
	}
	if count > 0 {
		getSql = "SELECT SUM(numberofcopies), SUM(numberofcopiesavailable), SUM(numberofcopiesborrowed) FROM booklist"
	} else {
		getSql = "SELECT 0, 0, 0 FROM DUAL"
	}
	rows, err := db.Query("SELECT Title FROM booklist")
	if err != nil {
		neem.DBError("error executing query", err)
		return types.LibraryData{}, fmt.Errorf("error in database")
	}
	defer rows.Close()
	var bookNames []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			neem.DBError("error scanning row", err)
			return types.LibraryData{}, fmt.Errorf("error in database")
		}
		bookNames = append(bookNames, name)
	}

	var data types.LibraryData

	data.Books = bookNames
	err = db.QueryRow(getSql).Scan(&data.NumberofCopies, &data.NumberofCopiesAvailable, &data.NumberofCopiesBorrowed)
	if err != nil {
		if err == sql.ErrNoRows {
			neem.Log("Books not present in database")
			return types.LibraryData{}, fmt.Errorf("no books in database")
		} else {
			neem.DBError("error getting from database", err)
			return types.LibraryData{}, fmt.Errorf("error in database")
		}
	} 
		
	return data, nil

}

func BooksList() ([]types.Book, error) {
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return []types.Book{}, err
	}
	rows, err := db.Query("SELECT * FROM booklist LIMIT 18")
	if err != nil {
		neem.DBError("error executing query", err)
		return []types.Book{}, err
	}
	defer rows.Close()
	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.B_Id, &book.Title, &book.Author, &book.Genre, &book.NumberofCopies, &book.NumberofCopiesAvailable, &book.NumberofCopiesBorrowed, &book.DueTime)
		if err != nil {
			neem.DBError("error scanning row", err)
			return []types.Book{}, err
		}
		books = append(books, book)
	}
	return books, nil

}

func GetCheckRequests() (types.AdminData, error){
	const query1 = `SELECT username FROM convertq;`
    const query2 = `SELECT booklist.Title, transactions.* FROM transactions INNER JOIN booklist ON transactions.B_Id = booklist.B_Id WHERE CheckOutAccepted IS NULL;`;
    const query3 = `SELECT booklist.Title, transactions.* FROM transactions INNER JOIN booklist ON transactions.B_Id = booklist.B_Id WHERE CheckInAccepted = 0 AND CheckOutAccepted = 1;`;
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return types.AdminData{}, err
	}
	rows, err := db.Query(query1)
	if err != nil {
		neem.DBError("error executing query", err)
		return types.AdminData{}, err
	}
	defer rows.Close()
	CliReq := []string{}
	var UserName string
	for rows.Next() {
		err := rows.Scan(&UserName)
		if err != nil {
			neem.DBError("error scanning row", err)
			return types.AdminData{}, err
		}
		CliReq = append(CliReq, UserName)
	}
	rows, err = db.Query(query2)
	if err != nil {
		neem.DBError("error executing query", err)
		return types.AdminData{}, err
	}
	defer rows.Close()
	CheckOutApprovals := []types.Transactions{}
	for rows.Next() {
		CheckOutApproval := types.Transactions{}
		var conv1, conv2 string
		err := rows.Scan(&CheckOutApproval.Title, &conv1, &CheckOutApproval.UserName, &conv2, &CheckOutApproval.CheckOutAccepted, &CheckOutApproval.CheckOutAccepted, &CheckOutApproval.DateBorrowed, &CheckOutApproval.DateReturned, &CheckOutApproval.OverdueFine)		
		if err != nil {
			neem.DBError("error scanning row", err)
			return types.AdminData{}, err
		}
		CheckOutApproval.B_Id, err = strconv.Atoi(conv1)
		if err != nil {
			neem.DBError("error converting string to int", err)
			return types.AdminData{}, err
		}
		CheckOutApproval.T_Id, err = strconv.Atoi(conv2)
		if err != nil {
			neem.DBError("error converting string to int", err)
			return types.AdminData{}, err
		}
		CheckOutApprovals = append(CheckOutApprovals, CheckOutApproval)
	}
	rows, err = db.Query(query3)
	if err != nil {
		neem.DBError("error executing query", err)
		return types.AdminData{}, err
	}
	defer rows.Close()
	CheckInApprovals := []types.Transactions{}
	for rows.Next() {
		CheckInApproval := types.Transactions{}
		var conv1, conv2 string
		err := rows.Scan(&CheckInApproval.Title, conv1, &CheckInApproval.UserName, conv2, &CheckInApproval.CheckOutAccepted, &CheckInApproval.CheckOutAccepted, &CheckInApproval.DateBorrowed, &CheckInApproval.DateReturned, &CheckInApproval.OverdueFine)		
		if err != nil {
			neem.DBError("error scanning row", err)
			return types.AdminData{}, err
		}
		CheckInApproval.B_Id, err = strconv.Atoi(conv1)
		if err != nil {
			neem.DBError("error converting string to int", err)
			return types.AdminData{}, err
		}
		CheckInApproval.T_Id, err = strconv.Atoi(conv2)
		if err != nil {
			neem.DBError("error converting string to int", err)
			return types.AdminData{}, err
		}
		CheckInApprovals = append(CheckInApprovals, CheckInApproval)
	}
	return types.AdminData{ConvertRequestClients: CliReq, CheckInApprovals: CheckInApprovals , CheckOutApprovals: CheckOutApprovals}, nil
}

func GetBook(title string) (types.Book, error) {
	db, err := Connection()
	if err != nil {
		neem.Critial(err, "error connecting to the database")
		return types.Book{}, err
	}
	rows, err := db.Query("SELECT * FROM booklist WHERE Title = ?", title)
	if err != nil {
		neem.DBError("error executing query", err)
		return types.Book{}, err
	}
	defer rows.Close()
	var book types.Book
	for rows.Next() {
		err := rows.Scan(&book.B_Id, &book.Title, &book.Author, &book.Genre, &book.NumberofCopies, &book.NumberofCopiesAvailable, &book.NumberofCopiesBorrowed, &book.DueTime)
		if err != nil {
			neem.DBError("error scanning row", err)
			return types.Book{}, err
		}
	}
	return book, nil
}