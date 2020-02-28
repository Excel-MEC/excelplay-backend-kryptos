package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	router := mux.NewRouter()
	logger := logrus.New()

	server := &http.Server{
		Handler:      newServer(router),
		Addr:         PORT,
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	logger.Println("Server starting on port " + PORT)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatalln("Could not start server on port " + PORT)
	} else {
		logger.Fatalln(err)
	}
}
