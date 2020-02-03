package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	if err := startServer(); err != nil {
		log.Fatal(err)
	}
}

func startServer() error {

	server := &http.Server{
		Handler:      getRouter(),
		Addr:         ":8000",
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}
	err := server.ListenAndServe()
	return err
}
