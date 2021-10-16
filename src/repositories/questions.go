package repositories

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"white-hat.api/src/entities"
)

var ctx = context.TODO()

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Questions struct {
	database *mongo.Database
}

func NewQuestions(database *mongo.Database) *Questions {
	return &Questions{database}
}

func (this *Questions) GetAll() []entities.Question {
	cursor, err := this.database.Collection("questions").Find(ctx, bson.D{{}})
	fatal(err)

	var questions []entities.Question

	for cursor.Next(ctx) {
		var question entities.Question

		err := cursor.Decode(&question)
		fatal(err)

		questions = append(questions, question)
	}

	err = cursor.Err()
	fatal(err)

	cursor.Close(ctx)

	return questions
}
