package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes(r *mux.Router) *mux.Router {
	for _, route := range s.getRoutes() {
		r.HandleFunc(route.path, route.handler).Methods(route.method)
	}

	return r
}

type route struct {
	path    string
	handler http.HandlerFunc
	method  string
}

func (s *Server) getRoutes() []*route {
	return []*route{
		{
			path:    "/users/{username}",
			handler: s.GetUser,
			method:  http.MethodGet,
		},
		{
			path:    "/users",
			handler: s.CreateUser,
			method:  http.MethodPost,
		},
	}
}
