package models

import (
	"github.com/Kamva/mgm/v2"
)

// Todo is the model of data of our database
// it inherits from MGM's DefaultModel interface
// wich allowed us to use methods and properties
// from mongoDB like _id or created_at fields
type Todo struct {
	mgm.DefaultModel `bson:",inline"`
	UserID           string `json:"userID" bson:"userID"`
	Title            string `json:"title" bson:"title"`
	Description      string `json:"description" bson:"description"`
	Done             bool   `json:"done" bson:"done"`
}

// CreateTodo is a wrapper function that creates a newTodo Object.
// It returns a Todo pointer to a Todo struct that contains all
// all the specified data
// & indicates that the returning value is the memory address for this
// created Todo
func CreateTodo(title, description, userID string) *Todo {
	return &Todo{
		UserID:      userID,
		Title:       title,
		Description: description,
		Done:        false,
	}
}
