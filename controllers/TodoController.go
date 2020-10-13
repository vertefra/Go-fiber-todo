package controllers

// in the controllers we define the functions tha we can use as fallbacks
// in the Fiber route

// defining the imports

import (
	"gitHub.com/vertefra/gofiber-todo-api/models"
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

// Every controller is a callback function that we pass as a second parameter

// GetAllTodos - GET /api/todos
// to our route definitions
func GetAllTodos(ctx *fiber.Ctx) {

	// SimpleFind is a function that takes two arguments:
	//
	// 1- 	memory address of the datastructure where the result
	//		should be stored (the todos that we created)
	// 2-	A filter. If filter is empty returns all the entries

	collection := mgm.Coll(&models.Todo{})
	todos := []models.Todo{}
	err := collection.SimpleFind(&todos, bson.D{})
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

}

// GetOneTodo - GET /api/todos/:id
func GetOneTodo(ctx *fiber.Ctx) {

}
