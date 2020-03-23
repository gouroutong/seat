package model

type User struct {
	Id int64
	Username string `json:"username"`
	Password string `json:"password"`
}
