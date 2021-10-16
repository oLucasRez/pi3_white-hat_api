package rests

import (
	"github.com/labstack/echo/v4"

	"white-hat.api/src/repositories"
	"white-hat.api/src/rests/server"
)

type dependencies struct {
	server *server.Server
	repo   *repositories.Questions
}

func (this *dependencies) getQuestions(web echo.Context) error {
	questions := this.repo.GetAll()

	return web.JSON(200, questions)
}

func RegisterQuestionsRoutes(server *server.Server, repo *repositories.Questions) {
	this := &dependencies{server, repo}

	this.server.Get("/questions", this.getQuestions)
}
