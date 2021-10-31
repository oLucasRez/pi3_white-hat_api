package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"white-hat.api/src/entities"
)

type Answers struct {
	database *mongo.Database
}

func NewAnswers(database *mongo.Database) *Answers {
	return &Answers{database}
}

func (this *Answers) Create(answers []entities.Answer, user_key string) {
	var obj interface{}
	obj = struct {
		user_key string
		answers  []entities.Answer
	}{user_key, answers}

	_, err := this.database.Collection("answers").InsertOne(ctx, obj)
	fatal(err)
}

func (this *Answers) GetAll() [][]entities.Answer {
	cursor, err := this.database.Collection("answers").Find(ctx, bson.D{{}})
	fatal(err)

	var answerss [][]entities.Answer

	for cursor.Next(ctx) {
		var answers []entities.Answer

		err := cursor.Decode(&answers)
		fatal(err)

		answerss = append(answerss, answers)
	}

	err = cursor.Err()
	fatal(err)

	cursor.Close(ctx)

	return answerss
}
