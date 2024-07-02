package types

import (

	"database/sql"
)

type UserLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	IsAdmin bool `json:"isadmin"`
}

type User struct {
	UserName string `json:"username"`
	HashedPassword string `json:"hashedpassword"`
	Salt string `json:"salt"`
	IsAdmin bool `json:"isadmin"`
}

type UserRegister struct {
	UserName string `json:"username"`
	HashedPassword string `json:"hashedpassword"`
	Salt string `json:"salt"`
	IsAdmin bool `json:"isadmin"`
}

type CookieUser struct {
	UserName string `json:"username"`
	IsAdmin bool `json:"isadmin"`
}

//TODO: Try eliminating this struct
type BookTemp struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Genre string `json:"genre"`
	DueTime string `json:"duetime"`
	Quantity string `json:"quantity"`
}

type Book struct {
	B_Id int `json:"b_id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Genre string `json:"genre"`
	DueTime int `json:"duetime"`
	Quantity int `json:"quantity"`
	NumberofCopies int `json:"numberofcopies"`
	NumberofCopiesAvailable int `json:"numberofcopiesavailable"`
	NumberofCopiesBorrowed int `json:"numberofcopiesborrowed"`
}

type Message struct {
	Message string `json:"message"`
}

type LibraryData struct {
	Books []string `json:"books"`
	NumberofCopies int `json:"numberofcopies"`
	NumberofCopiesAvailable int `json:"numberofcopiesavailable"`
	NumberofCopiesBorrowed int `json:"numberofcopiesborrowed"`
}

type PageDataAdmin struct {
    LibraryData  LibraryData
	AdminData AdminData
}

type PageDataClient struct {
    LibraryData  LibraryData
	Books []Book
}

type Transactions struct {
	Title string `json:"title"`
	T_Id int `json:"t_id"`
	B_Id int `json:"b_id"`
	UserName string `json:"username"`
	CheckInAccepted sql.NullBool `json:"checkinaccepted"`
	CheckOutAccepted sql.NullBool `json:"checkoutaccepted"`
	DateBorrowed sql.NullTime `json:"dateborrowed"`
	DateReturned sql.NullTime `json:"datereturned"`
	DueTime int `json:"duetime"`
	OverdueFine int `json:"overduefine"`
}

type AdminData struct {
	ConvertRequestClients []string `json:"convertrequestclients"`
	CheckInApprovals []Transactions `json:"checkinapprovals"`
	CheckOutApprovals []Transactions `json:"checkoutapprovals"`
}