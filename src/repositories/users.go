package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"white-hat.api/src/entities"
)

type Users struct {
	database *mongo.Database
}

func NewUsers(database *mongo.Database) *Users {
	return &Users{database}
}

func (this *Users) Create(user entities.User) {
	_, err := this.database.Collection("users").InsertOne(ctx, user)
	fatal(err)
}
