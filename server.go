package main

import (
	"log"
	"os"

	"gitHub.com/vertefra/gofiber-todo-api/controllers"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var key string

// Init starts the conncetion with mongoDb
func init() {
	godotenv.Load()
	var mongoURI string
	mongoDb := os.Getenv("MONGO_DB")
	mongoUsr := os.Getenv("MONGO_USR")
	mongoPsw := os.Getenv("MONGO_PSW")
	key = os.Getenv("SECRET")

	if len(mongoDb) == 0 || len(mongoUsr) == 0 || len(mongoPsw) == 0 {
		mongoURI = "mongodb://localhost:27017"
	} else {
		mongoURI = "mongodb+srv://" + mongoUsr + ":" + mongoPsw + "@cluster0-fg0dv.gcp.mongodb.net/" + mongoDb + "?retryWrites=true&w=majority"
	}

	if err := mgm.SetDefaultConfig(nil, mongoDb, options.Client().ApplyURI(mongoURI)); err != nil {
		log.Fatal(err)
	}

	log.Printf("Server Connected to %s", mongoDb)
}

func main() {

	app := fiber.New()

	app.Post("/api/users/login", controllers.Login)
	// app.Post("/api/users/signup")

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(key),
	}))

	app.Get("/api/todos", controllers.GetAllTodos)
	app.Get("/api/todos/:id", controllers.GetTodoByID)
	app.Post("/api/todos", controllers.CreateTodo)
	app.Patch("/api/todos/:id", controllers.UpdateTodo)
	app.Delete("/api/todos/:id", controllers.DeleteTodo)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal("error in Listen on 3000")
	} else {
		log.Println("Server Listening on port 3000")
	}
}
