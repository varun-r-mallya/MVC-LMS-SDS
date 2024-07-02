package controllers

import (
	"net/http"
	"encoding/json"
	"io"
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
		Title: bookTemp.Title,
		Genre: bookTemp.Genre,
		Author: bookTemp.Author,
		DueTime: duetime,
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

