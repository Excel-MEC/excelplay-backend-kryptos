package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type server struct {
	router *mux.Router
	db     *sqlx.DB
}

func newServer(router *mux.Router, db *sqlx.DB) *server {
	server := &server{
		router,
		db,
	}
	server.routes() // Register the route handling functions with the mux router
	return server
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
