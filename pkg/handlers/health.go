package handlers

import (
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
)

// HeartBeat sends a response back, just to check if the server is up.
// @Summary ping example
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /api/ping [get]
func HeartBeat() httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test"))
		return nil
	}
}
