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

func (s *server) routes() {
	s.router.HandleFunc("/admin/", s.handleAdmin).Methods("GET")
}

func (s *server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin"))
}
