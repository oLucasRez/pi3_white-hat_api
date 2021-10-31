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
	obj := struct {
		User_key string            `bson:"user_key"`
		Answers  []entities.Answer `bson:"answers"`
	}{user_key, answers}

	_, err := this.database.Collection("answers").InsertOne(ctx, obj)
	fatal(err)
}

func (this *Answers) GetByUserKey(user_key string) [][]entities.Answer {
	cursor, err := this.database.Collection("answers").Find(ctx, bson.M{"user_key": user_key})
	fatal(err)

	var answerss [][]entities.Answer

	for cursor.Next(ctx) {
		var answers struct {
			User_key string            `bson:"user_key"`
			Answers  []entities.Answer `bson:"answers"`
		}

		err := cursor.Decode(&answers)
		fatal(err)

		answerss = append(answerss, answers.Answers)
	}

	err = cursor.Err()
	fatal(err)

	cursor.Close(ctx)

	return answerss
}
