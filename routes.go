package main

import (
	"github.com/Excel-MEC/excelplay-backend-kryptos/controllers"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api", controllers.CheckAPIIsAlive).Methods("GET")

	return router
}
