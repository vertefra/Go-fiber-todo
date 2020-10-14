package controllers

import (
	"gitHub.com/vertefra/gofiber-todo-api/models"
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllUsers - GET /api/users
// Returns all the users currently in the database
func GetAllUsers(ctx *fiber.Ctx) {

	collection := mgm.Coll(&models.User{})
	users := []models.User{}

	if err := collection.SimpleFind(&users, bson.D{}); err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":    true,
		"users": users,
	})
}
