package entity

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

type UserID int

func NewUser(email string, password string) User {
	return User{
		Password: password,
		Email:    email,
	}
}

type User struct {
	ID           UserID
	Email        string
	Password     string
	UserAtribute UserAttribute
}

func (u User) EncryptPassword() string {
	converted, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	return hex.EncodeToString(converted[:])
}

type UserAttribute struct {
	UserId UserID
	Name   string
	Sex    string
}
