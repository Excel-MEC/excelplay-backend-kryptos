package main

import (
	"excelplay-backend/controllers"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api", controllers.CheckAPIIsAlive).Methods("GET")

	return router
}
