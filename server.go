package main

import (
	"excelplay-backend/logger"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Handler:      getRouter(),
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
