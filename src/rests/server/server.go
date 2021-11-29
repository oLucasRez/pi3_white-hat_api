package server

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func New() *Server {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	return &Server{e}
}

func (server *Server) Run() {
	port_string := ":"
	port := os.Getenv("PORT")

	if port != "" {
		port_string += port
	} else {
		port_string += "1323"
	}

	server.echo.Logger.Fatal(server.echo.Start(port_string))
}

func (server *Server) Get(path string, handler echo.HandlerFunc) {
	server.echo.GET(path, handler)
}

func (server *Server) Post(path string, handler echo.HandlerFunc) {
	server.echo.POST(path, handler)
}

func (server *Server) Put(path string, handler echo.HandlerFunc) {
	server.echo.PUT(path, handler)
}

func (server *Server) Delete(path string, handler echo.HandlerFunc) {
	server.echo.DELETE(path, handler)
}
