package server

import (
	"github.com/iancharters/gostart/internal/app"
)

type Server struct {
	app *app.App
}

func New(app *app.App) *Server {
	return &Server{
		app: app,
	}
}
