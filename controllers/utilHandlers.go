package controllers

import "net/http"

// CheckAPIIsAlive sends a response back, just to check if the server is up.
func CheckAPIIsAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Test"))
}
