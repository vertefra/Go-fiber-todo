package controllers

// in the controllers we define the functions tha we can use as fallbacks
// in the Fiber route

// defining the imports

import (
	"log"

	"gitHub.com/vertefra/gofiber-todo-api/models"
	"github.com/Kamva/mgm/v2"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

// Every controller is a callback function that we pass as a second parameter

// GetAllTodos - GET /api/todos?user=id
// to our route definitions
func GetAllTodos(ctx *fiber.Ctx) {

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	log.Println("claims --> ", email)

	userID := ctx.Query("user")

	collection := mgm.Coll(&models.Todo{})
	todos := []models.Todo{}

	err := collection.SimpleFind(&todos, bson.M{"UserID": userID})

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":    true,
		"todos": todos,
	})

}

// GetTodoByID - GET /api/todos/:id
// Creates a pointer to our Todo struct
// and use that pointer to define the structure
// of our collection
func GetTodoByID(ctx *fiber.Ctx) {

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	log.Println("claims --> ", email)

	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

// CreateTodo - POST /api/todos?user=id
func CreateTodo(ctx *fiber.Ctx) {

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	log.Println("claims --> ", email)

	body := new(struct {
		Title       string
		Description string
	})

	ctx.BodyParser(&body)

	userID := ctx.Query("user")

	if len(body.Title) == 0 || len(body.Description) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Empty fields",
		})
		return
	}

	if len(userID) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "no user id found",
		})
		return
	}

	// TODO: check if the user id is valid

	//

	todo := &models.Todo{
		UserID:      userID,
		Title:       body.Title,
		Description: body.Description,
	}

	todo = models.CreateTodo(todo.Title, todo.Description, todo.UserID)
	err := mgm.Coll(todo).Create(todo)

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})

		return
	}

	ctx.Status(201).JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

// UpdateTodo PATCH - PATCH /api/todos/:id
func UpdateTodo(ctx *fiber.Ctx) {

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	log.Println("claims --> ", email)

	id := ctx.Params("id")

	body := new(struct {
		Title       string
		Description string
		Done        bool
	})

	ctx.BodyParser(&body)

	todo := &models.Todo{}

	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	if len(body.Title) > 0 || len(body.Description) > 0 || body.Done != todo.Done {

		todo.Title = body.Title
		todo.Description = body.Description
		todo.Done = body.Done

		collection.Update(todo)

		ctx.Status(200).JSON(fiber.Map{
			"ok":      true,
			"updated": todo,
		})

	}

}

// DeleteTodo -  DELETE /api/todos/:id
func DeleteTodo(ctx *fiber.Ctx) {

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	log.Println("claims --> ", email)

	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	if err := collection.FindByID(id, todo); err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	if err := collection.Delete(todo); err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":      true,
		"deleted": todo,
	})
}
