package main

import (
	"log"
	"os"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Init starts the conncetion with mongoDb
func initMongo() {
	godotenv.Load()
	var mongoURI string
	mongoDb := os.Getenv("MONGO_DB")
	mongoUsr := os.Getenv("MONGO_USR")
	mongoPsw := os.Getenv("MONGO_PSW")

	if len(mongoDb) == 0 || len(mongoUsr) == 0 || len(mongoPsw) == 0 {
		mongoURI = "mongodb://localhost:27017"
	} else {
		mongoURI = "mongodb+srv://" + mongoUsr + ":" + mongoPsw + "@cluster0-fg0dv.gcp.mongodb.net/" + mongoDb + "?retryWrites=true&w=majority"
	}

	if err := mgm.SetDefaultConfig(nil, mongoDb, options.Client().ApplyURI(mongoURI)); err != nil {
		log.Fatal(err)
	}
}

func main() {

	app := fiber.New()

}
