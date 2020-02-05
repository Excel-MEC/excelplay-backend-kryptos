package main

import (
	"excelplay-backend-kryptos/src/controllers"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api", controllers.CheckAPIIsAlive).Methods("GET")

	return router
}
