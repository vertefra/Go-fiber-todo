package controllers

import (
	"fmt"

	"gitHub.com/vertefra/gofiber-todo-api/models"
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllUsers - GET /api/users
// Returns all the users currently in the database
func GetAllUsers(c *fiber.Ctx) {

	collection := mgm.Coll(&models.User{})
	users := []models.User{}

	if err := collection.SimpleFind(&users, bson.D{}); err != nil {
		c.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	c.Status(200).JSON(fiber.Map{
		"ok":    true,
		"users": users,
	})
}

// AddNewUser - POST /api/users/signup
// create a New User
func AddNewUser(ctx *fiber.Ctx) {

	params := new(struct {
		Email    string
		Password string
	})

	fmt.Println(params)

}
