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

// docker tag white-hat-api registry.digitalocean.com/white-hat-api/white-hat-api:0.0.2
// docker push registry.digitalocean.com/white-hat-api/white-hat-api:0.0.2

var ctx = context.TODO()

func main() {
	database := connectMongo()

	migrations.InsertQuestions(database)

	server := server.New()

	users_repo := repositories.NewUsers(database)
	rests.RegisterUsersRoutes(server, users_repo)

	questions_repo := repositories.NewQuestions(database)
	rests.RegisterQuestionsRoutes(server, questions_repo)

	answers_repo := repositories.NewAnswers(database)
	rests.RegisterAnswersRoutes(server, answers_repo, questions_repo)

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
