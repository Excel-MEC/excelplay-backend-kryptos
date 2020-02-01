package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	router := mux.NewRouter()

	server := newServer(router)
	log.Fatal(http.ListenAndServe(":8001", server))
	return nil
}
