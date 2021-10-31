package rests

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"white-hat.api/src/entities"
	"white-hat.api/src/repositories"
	"white-hat.api/src/rests/server"
)

type answers_deps struct {
	server         *server.Server
	answers_repo   *repositories.Answers
	questions_repo *repositories.Questions
}

func (this *answers_deps) postAnswers(web echo.Context) error {
	user_key := web.Param("key")

	var answers []entities.Answer
	err := json.NewDecoder(web.Request().Body).Decode(&answers)
	fatal(err)

	result := this.calculateResult(answers)

	this.answers_repo.Create(answers, user_key)

	return web.JSON(200, result)
}

func (this *answers_deps) getResults(web echo.Context) error {
	user_key := web.Param("key")

	answerss := this.answers_repo.GetByUserKey(user_key)

	results := []entities.Result{}
	for _, answers := range answerss {
		results = append(results, this.calculateResult(answers))
	}

	return web.JSON(200, results)
}

func (this *answers_deps) calculateResult(answers []entities.Answer) entities.Result {
	questions := this.questions_repo.GetAll()

	var result entities.Result = make(entities.Result)
	for _, answer := range answers {
		focus := questions[answer.Id].Focus
		category := questions[answer.Id].Category

		if result[focus] == nil {
			result[focus] = make(map[string]int)
		}

		result[focus][category] += answer.Rate
	}

	return result
}

func RegisterAnswersRoutes(server *server.Server, answers_repo *repositories.Answers, questions_repo *repositories.Questions) {
	this := &answers_deps{server, answers_repo, questions_repo}

	server.Post("/answers/:key", this.postAnswers)
	server.Get("/answers/:key", this.getResults)
}
