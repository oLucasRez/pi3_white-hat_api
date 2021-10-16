package rests

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"white-hat.api/src/entities"
	"white-hat.api/src/rests/server"
)

func postAnswer(c echo.Context) error {
	id, _ := strconv.ParseInt(c.FormValue("id"), 10, 64)
	rate, _ := strconv.ParseInt(c.FormValue("rate"), 10, 64)

	answers := &entities.Answer{Id: int(id), Rate: int(rate)}

	return c.JSON(200, answers)
}

func RegisterAnswersRoutes(server *server.Server) {
	server.Post("/answers", postAnswer)
}
