package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func newServer(router *mux.Router) *server {
	server := &server{
		router: router,
	}
	server.routes()
	return server
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
