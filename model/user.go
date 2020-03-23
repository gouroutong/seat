package model

import (
	"errors"
	"fmt"
)

type User struct {
	Id int64
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Login(user1 *User) error {
	if DB.Where("username = ?", user.Username).First(user1).RecordNotFound() {
		return errors.New("username does not exist")
	}
	fmt.Println("user", user)
	if user.Password != user1.Password {
		return errors.New("incorrect password")
	}
	return nil
}