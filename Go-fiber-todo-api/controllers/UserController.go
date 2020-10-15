package controllers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"gitHub.com/vertefra/gofiber-todo-api/models"
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

func authenticate(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	key := os.Getenv("SECRET")

	// setting claims

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Generate encoded token

	t, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return t, nil
}

// Signup - POST /api/users/signup
// Returns all the users currently in the database
func Signup(ctx *fiber.Ctx) {

	body := new(struct {
		Email    string
		Password string
	})

	ctx.BodyParser(&body)

	user := &models.User{}

	if foundUser := mgm.Coll(user).FindOne(mgm.Ctx(), bson.M{"email": body.Email}).Decode(&user); foundUser != nil {

		// The user does not exist, create the user

		fmt.Println(&foundUser)

		user := models.CreateUser(body.Email, body.Password)

		if err := mgm.Coll(user).Create(user); err != nil {
			ctx.Status(400).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
			return
		}

		// the user has been created, let's authenticate

		t, err := authenticate(body.Email)

		if err != nil {
			ctx.Status(404).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}

		ctx.Status(200).JSON(fiber.Map{
			"ok":    true,
			"token": t,
		})

	}

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
		return
	}

	t, err := authenticate(body.Email)

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":    true,
		"token": t,
	})

}
