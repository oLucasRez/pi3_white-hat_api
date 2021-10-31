package rests

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"white-hat.api/src/entities"
	"white-hat.api/src/repositories"
	"white-hat.api/src/rests/server"
)

type users_deps struct {
	server *server.Server
	repo   *repositories.Users
}

func (this *users_deps) postUsers(web echo.Context) error {
	users := entities.User{Key: uuid.New().String()}

	this.repo.Create(users)

	return web.JSON(200, users)
}

func RegisterUsersRoutes(server *server.Server, repo *repositories.Users) {
	this := &users_deps{server, repo}

	this.server.Post("/users", this.postUsers)
}
