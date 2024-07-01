package types

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

type Message struct {
	Message string `json:"message"`
}