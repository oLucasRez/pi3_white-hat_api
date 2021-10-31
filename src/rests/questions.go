package rests

import (
	"github.com/labstack/echo/v4"

	"white-hat.api/src/repositories"
	"white-hat.api/src/rests/server"
)

type questions_deps struct {
	server *server.Server
	repo   *repositories.Questions
}

func (this *questions_deps) getQuestions(web echo.Context) error {
	questions := this.repo.GetAll()

	return web.JSON(200, questions)
}

func RegisterQuestionsRoutes(server *server.Server, repo *repositories.Questions) {
	this := &questions_deps{server, repo}

	this.server.Get("/questions", this.getQuestions)
}
