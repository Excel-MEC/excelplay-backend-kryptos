package handlers

import (
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
)

// HeartBeat sends a response back, just to check if the server is up.
func HeartBeat() httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test"))
		return nil
	}
}
