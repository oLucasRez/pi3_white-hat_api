package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func New() *Server {
	e := echo.New()

	e.Use(middleware.CORS())

	return &Server{e}
}

func (server *Server) Run() {
	server.echo.Logger.Fatal(server.echo.Start(":1323"))
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
