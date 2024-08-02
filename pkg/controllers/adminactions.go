package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/models"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"
)

func AddBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Admin add books API accessed")
	var bookTemp types.BookTemp
	body, err := io.ReadAll(r.Body)
	if err != nil {
		neem.Spotlight(err, "could not read body in addbooks")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &bookTemp)
	if err != nil {
		neem.Spotlight(err, "could not unmarshal book")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	duetime, err := strconv.Atoi(bookTemp.DueTime)
	if err != nil {
		neem.Log("could not convert DueTime to int")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	quantity, err := strconv.Atoi(bookTemp.Quantity)
	if err != nil {
		neem.Log("could not convert DueTime to int")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	book := types.Book{
		Title:    bookTemp.Title,
		Genre:    bookTemp.Genre,
		Author:   bookTemp.Author,
		DueTime:  duetime,
		Quantity: quantity,
	}

	if book.Quantity <= 0 {
		neem.Log("Quantity less than 0")
		toSend := types.Message{Message: "Quantity cannot be less than 0"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return

	}
	if book.DueTime <= 0 {
		neem.Log("Quantity less than 0")
		toSend := types.Message{Message: "Due time cannot be less than 0"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return

	}
	success, err := models.AddBooks(book)
	if err != nil {
		neem.Log("could not add book")
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusInternalServerError)
		return
	}

	if success {
		neem.Log("Book added successfully")
		toSend := types.Message{Message: "Book Added Successfully"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusOK)
		return
	}
}

func UpdateBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Admin update books API accessed")
	var bookTemp types.BookTemp
	body, err := io.ReadAll(r.Body)
	if err != nil {
		neem.Spotlight(err, "could not read body in updatebooks")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &bookTemp)
	if err != nil {
		neem.Spotlight(err, "could not unmarshal book")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	duetime, err := strconv.Atoi(bookTemp.DueTime)
	if err != nil {
		neem.Log("could not convert DueTime to int")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	quantity, err := strconv.Atoi(bookTemp.Quantity)
	if err != nil {
		neem.Log("could not convert DueTime to int")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	bookID, err := strconv.Atoi(bookTemp.B_Id)
	if err != nil {
		neem.Log("could not convert bookID to int")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	book := types.Book{
		Title:    bookTemp.Title,
		Genre:    bookTemp.Genre,
		Author:   bookTemp.Author,
		DueTime:  duetime,
		Quantity: quantity,
		B_Id:     bookID,
	}

	if book.Quantity <= 0 {
		neem.Log("Quantity less than 0")
		toSend := types.Message{Message: "Quantity cannot be less than 0"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return

	}
	if book.DueTime <= 0 {
		neem.Log("Quantity less than 0")
		toSend := types.Message{Message: "Due time cannot be less than 0"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return

	}
	success, err := models.UpdateBooks(book)
	if err != nil {
		neem.Log("could not update book")
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusInternalServerError)
		return
	}

	if success {
		neem.Log("Book Updated successfully")
		toSend := types.Message{Message: "Book Updated Successfully"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusOK)
		return
	}
}

func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Admin delete books API accessed")
	var bookTemp types.BookTemp
	body, err := io.ReadAll(r.Body)
	if err != nil {
		neem.Spotlight(err, "could not read body in deletebooks")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &bookTemp)
	if err != nil {
		neem.Spotlight(err, "could not unmarshal book")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	bookID, err := strconv.Atoi(bookTemp.B_Id)
	if err != nil {
		neem.Log("could not convert bookID to int")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	success, err := models.DeleteBooks(bookID)
	if err != nil {
		neem.Log("could not delete book")
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		w.Write(b)
		return
	}

	if success {
		neem.Log("Book deleted successfully")
		toSend := types.Message{Message: "Book Deleted Successfully"}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		w.Write(b)
		return
	}
}

func AcceptCheckOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Admin accept checkout API accessed")
	var checkOut types.CheckOut
	body, err := io.ReadAll(r.Body)
	if err != nil {
		neem.Spotlight(err, "could not read body in acceptcheckout")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &checkOut)
	if err != nil {
		neem.Spotlight(err, "could not unmarshal checkout")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	message, err := models.AcceptCheckOut(checkOut)
	if err != nil {
		neem.Log("could not accept checkout")
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusInternalServerError)
		return
	}
	neem.Log("Checkout processed successfully")
	toSend := types.Message{Message: message}
	fmt.Println(message)
	b, err := json.Marshal(toSend)
	if err != nil {
		neem.Spotlight(err, "could not marshal message")
	}
	w.Write(b)
}

func AcceptCheckIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Admin accept checkin API accessed")
	var checkIn types.CheckIn
	body, err := io.ReadAll(r.Body)
	if err != nil {
		neem.Spotlight(err, "could not read body in acceptcheckin")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &checkIn)
	if err != nil {
		neem.Spotlight(err, "could not unmarshal checkIn")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	message, err := models.AcceptCheckIn(checkIn)
	if err != nil {
		neem.Log("could not accept checkIn")
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusInternalServerError)
		return
	}
	neem.Log("CheckIn processed successfully")
	toSend := types.Message{Message: message}
	fmt.Println(message)
	b, err := json.Marshal(toSend)
	if err != nil {
		neem.Spotlight(err, "could not marshal message")
	}
	w.Write(b)
}

func AcceptAdmins(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	neem.Log("Admin accept API accessed")
	var AcceptAdmins types.AcceptAdmins
	body, err := io.ReadAll(r.Body)
	if err != nil {
		neem.Spotlight(err, "could not read body in acceptadmin")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &AcceptAdmins)
	if err != nil {
		neem.Spotlight(err, "could not unmarshal acceptadmin")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	message, err := models.AcceptAdmins(AcceptAdmins)
	if err != nil {
		neem.Log("could not accept admin request")
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			neem.Spotlight(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusInternalServerError)
		return
	}
	neem.Log("AcceptAdmins processed successfully")
	toSend := types.Message{Message: message}
	fmt.Println(message)
	b, err := json.Marshal(toSend)
	if err != nil {
		neem.Spotlight(err, "could not marshal message")
	}
	w.Write(b)
}
