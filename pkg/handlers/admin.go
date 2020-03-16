package handlers

import "net/http"

// HandleAdmin handles any requests to the /api/admin endpoint
func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin"))
}
