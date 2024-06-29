package types

type UserLogin struct {
	UserName string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	IsAdmin bool `json:"isadmin"`
}

type UserRegister struct {
	UserName string `json:"username"`
	HashedPassword string `json:"hashedpassword"`
	Salt string `json:"salt"`
	IsAdmin bool `json:"isadmin"`
}


type Message struct {
	Message string `json:"message"`
}