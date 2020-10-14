package controllers

import (
	"fmt"
	"log"

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

// Login - POST /api/users/login
func Login(ctx *fiber.Ctx) {

	body := new(struct {
		Email    string
		Password string
	})

	ctx.BodyParser(&body)

	user := &models.User{}

	log.Println("Logging -> ", body.Email)

	collection := mgm.Coll(user)

	if len(body.Email) == 0 || len(body.Password) == 0 {
		log.Println("empty fields found")
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "empty fields",
		})
		return
	}

	if err := collection.SimpleFind(user, bson.M{"email": body.Email}); err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

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
