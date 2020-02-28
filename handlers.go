package main

import "net/http"

func (s *server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin"))
}

// CheckAPIIsAlive sends a response back, just to check if the server is up.
func (s *server) checkAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Test"))
}
