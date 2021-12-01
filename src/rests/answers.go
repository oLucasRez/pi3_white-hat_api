package rests

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"white-hat.api/src/entities"
	"white-hat.api/src/enums"
	"white-hat.api/src/repositories"
	"white-hat.api/src/rests/server"
)

type answers_deps struct {
	server         *server.Server
	answers_repo   *repositories.Answers
	questions_repo *repositories.Questions
}

func (bla *answers_deps) postAnswers(web echo.Context) error {
	user_key := web.Param("key")

	var answers []entities.Answer
	err := json.NewDecoder(web.Request().Body).Decode(&answers)
	fatal(err)

	result := bla.calculateResult(answers)

	bla.answers_repo.Create(answers, user_key)

	return web.JSON(200, result)
}

func (bla *answers_deps) getResults(web echo.Context) error {
	user_key := web.Param("key")

	answerss := bla.answers_repo.GetByUserKey(user_key)

	results := []entities.Result{}
	for _, answers := range answerss {
		results = append(results, bla.calculateResult(answers))
	}

	return web.JSON(200, results)
}

func (bla *answers_deps) calculateResult(answers []entities.Answer) entities.Result {
	questions := bla.questions_repo.GetAll()

	result := entities.Result{}
	normalizador_focus := map[string]int{}
	normalizador_category := map[string]int{}
	for _, answer := range answers {
		focus := questions[answer.Id-1].Focus
		category := questions[answer.Id-1].Category

		if result[focus].Result == nil {
			result[focus] = entities.Focuses{Result: map[string]float32{}}
		}

		result[focus].Result[category] += float32(answer.Rate)
		result[focus] = entities.Focuses{
			Result: result[focus].Result,
			Nota:   result[focus].Nota + float32(answer.Rate),
		}

		normalizador_focus[focus]++
		normalizador_category[category]++
	}

	for focus, focuses := range result {
		for category, _ := range focuses.Result {
			result[focus].Result[category] /= float32(normalizador_category[category])
		}

		Nota := result[focus].Nota / float32(normalizador_focus[focus])

		var Comment string
		if Nota >= 3 {
			Comment = enums.Comments[focus].Good
		} else {
			Comment = enums.Comments[focus].Bad
		}

		result[focus] = entities.Focuses{
			Comment: Comment,
			Nota:    Nota,
			Result:  focuses.Result,
		}
	}

	return result
}

func RegisterAnswersRoutes(server *server.Server, answers_repo *repositories.Answers, questions_repo *repositories.Questions) {
	this := &answers_deps{server, answers_repo, questions_repo}

	server.Post("/answers/:key", this.postAnswers)
	server.Get("/answers/:key", this.getResults)
}
