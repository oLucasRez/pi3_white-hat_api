package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"white-hat.api/src/migrations"
	"white-hat.api/src/repositories"
	"white-hat.api/src/rests"
	"white-hat.api/src/rests/server"
)

// docker-compose up --build

var ctx = context.TODO()

func main() {
	database := connectMongo()

	migrations.InsertQuestions(database)

	server := server.New()

	questions_repo := repositories.NewQuestions(database)

	rests.RegisterAnswersRoutes(server)
	rests.RegisterQuestionsRoutes(server, questions_repo)

	server.Run()
}

func connectMongo() *mongo.Database {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		panic(err)
	}

	err = client.Database("white-hat_db").Drop(ctx)

	if err != nil {
		panic(err)
	}

	return client.Database("white-hat_db")
}
