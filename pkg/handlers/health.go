package handlers

import (
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
)

// HeartBeat sends a response back, just to check if the server is up.
// @Summary Server health check.
// @Description Sends "Test" back. Use this to check if the server is up.
// @Tags Kryptos
// @Produce plain
// @Success 200 {string} string "Server is up"
// @Router /api/ping [get]
func HeartBeat() httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test"))
		return nil
	}
}
