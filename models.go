package main

import (
	"fmt"
	"time"
)



type UserManager struct {
	user		User
}


type User struct {
	Username	string
	Password	string
	Id		string
}


func (um UserManager) CreateUser(username string, password string) {
	user := new(user)

	user.Username = username
	user.Password = password

	um.user = user
	um.GenerateId()

	//write User model to database.
}

func (um UserManager) CheckIfCredentialsExist (username string, password string) bool {
		// check if credentials exists

		return false
}

func (um UserManager) CheckIfUsernameExist (username string) bool {
		// check if credentials exists

		return false
}

