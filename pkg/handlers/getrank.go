package handlers

import (
	"net/http"
	"strconv"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/liveleaderboard"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
	"github.com/dgrijalva/jwt-go"
)

// GetRank fetches the rank of a user from the in-memory leaderboard
// @Summary returns the rank of the user who made the request.
// @Description returns the rank of the user who made the request. Returns -1 for an invalid user ID.
// @Tags Kryptos
// @Produce plain
// @Success 200 {string} string "Returns the rank of the user who made the request. Returns -1 for an invalid user ID."
// @Failure 500 {string} string
// @Router /api/getrank [get]
func GetRank(db *database.DB, env *env.Config) httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		// Obtain values from JWT
		props, _ := r.Context().Value("props").(jwt.MapClaims)
		userID, _ := strconv.Atoi(props["user_id"].(string))

		liveleaderboard.FetchRank <- userID
		rank := <-liveleaderboard.ReturnRank

		// jsonRes, err := json.Marshal(res)
		// if err != nil {
		// 	return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
		// }
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(rank)))
		return nil
	}
}
