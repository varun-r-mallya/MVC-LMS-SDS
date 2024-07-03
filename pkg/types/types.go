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
	B_Id string `json:"bookID"`
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
type AdminData struct {
	ConvertRequestClients []string `json:"convertrequestclients"`
	CheckInApprovals []Transactions `json:"checkinapprovals"`
	CheckOutApprovals []Transactions `json:"checkoutapprovals"`
}

type PageDataAdmin struct {
    LibraryData  LibraryData
	AdminData AdminData
}

type PageDataClient struct {
    LibraryData  LibraryData
	Books []Book
	Transactions []ClientBookViewTransactionsInterpretable
}

//made changes to nullint16
type Transactions struct {
	Title string `json:"title"`
	T_Id int `json:"t_id"`
	B_Id int `json:"b_id"`
	UserName string `json:"username"`
	CheckInAccepted sql.NullBool `json:"checkinaccepted"`
	CheckOutAccepted sql.NullBool `json:"checkoutaccepted"`
	DateBorrowed []uint8 `json:"dateborrowed"`
	DateReturned []uint8 `json:"datereturned"`
	DueTime int `json:"duetime"`
	OverdueFine int `json:"overduefine"`
}

type ClientBookViewTransactions struct {
	Title string `json:"title"`
	UserName string `json:"username"`
	CheckInAccepted sql.NullBool `json:"checkinaccepted"`
	CheckOutAccepted sql.NullBool `json:"checkoutaccepted"`
	DateBorrowed []uint8 `json:"dateborrowed"`
	DateReturned []uint8 `json:"datereturned"`
	DueTime int `json:"duetime"`
	OverdueFine int `json:"overduefine"`
	Author string `json:"author"`
}

type ClientBookViewTransactionsInterpretable struct {
	Title string `json:"title"`
	UserName string `json:"username"`
	CheckInAccepted string `json:"checkinaccepted"`
	CheckOutAccepted string `json:"checkoutaccepted"`
	DateBorrowed string `json:"dateborrowed"`
	DateReturned string `json:"datereturned"`
	DueTime string `json:"duetime"`
	OverdueFine string `json:"overduefine"`
	Author string `json:"author"`
}
type ClientBookView struct {
	Book Book	`json:"book"`
	Transactions []ClientBookViewTransactionsInterpretable `json:"transactions"`
}

type AdminBookView struct {
	Book Book	`json:"book"`
	Transactions []ClientBookViewTransactionsInterpretable `json:"transactions"`
}

type CheckOut struct {
	T_Id int `json:"t_id"`
	Accepted bool `json:"accepted"`
}

type CheckIn struct {
	T_Id int `json:"t_id"`
	Accepted bool `json:"accepted"`
}

type AcceptAdmins struct {
	UserName string `json:"username"`
	Accepted bool `json:"accepted"`
}