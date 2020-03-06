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

func (s *server) handleNextQuestion() http.HandlerFunc {
	type response struct {
		Question   string `json:"question"`
		ImageLevel bool   `json:"image_level"`
		LevelFile  string `json:"level_file"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// do handling
	}
}
