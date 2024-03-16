package entities

import (
	"go-e-s/src/backbone/adapter"
	//"go-e-s/backbone/service_layer"
)

var _ adapter.IBaseEntity = &User{}

type User struct {
	// implement: adapter.IBaseEntity
	adapter.BaseEntity
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
}

func (u User) Create(firstName string, lastName string, username string, password string) *User {
	user := &User{
		FirstName: firstName,
		LastName:  lastName,
		UserName:  username,
		Password:  password,
	}
	return user
}

func (u User) Table() string {
	return "users"
}
