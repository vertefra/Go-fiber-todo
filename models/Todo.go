package models

import "github.com/Kamva/mgm/v2"

// Todo is the model of data of our database
// it inherits from MGM's DefaultModel interface
// wich allowed us to use methods and properties
// from mongoDB like _id or created_at fields
type Todo struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Description      string `json:"description" bson:"description"`
	Done             bool   `json:"done" bson:"done"`
}

// CreateTodo is a wrapper function that creates a newTodo Object.
// It returns a Todo pointer to a Todo struct that contains all
// all the specified data
func CreateTodo(title, description string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Done:        false,
	}
}
