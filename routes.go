package main

import (
	"github.com/Excel-MEC/excelplay-backend-kryptos/controllers"
)

func (s *server) routes() {
	s.router.HandleFunc("/admin/", s.handleAdmin).Methods("GET")
	s.router.HandleFunc("/api", s.checkAPIIsAlive).Methods("GET")
}
