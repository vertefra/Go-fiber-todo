package models

import "github.com/Kamva/mgm/v2"

// User struct. Store all the references to the todos, password and username
type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string   `json:"email" bson:"email"`
	Password         string   `json:"password" bson:"password"`
	Todos            []string `json:"todos" bson:"todos"`
}

// CreateUser returns a pointer to a User struct
func CreateUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}

}
