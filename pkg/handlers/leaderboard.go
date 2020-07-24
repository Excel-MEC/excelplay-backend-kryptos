package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
)

// GetLeaderboard returns all the users ordered in descending order of level,
// and for users on the same level, in the ascending order of last submission time.
func GetLeaderboard(db *database.DB, config *env.Config) httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		var users []database.User
		err := db.GetLeaderboard(&users)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Unable to fetch leaderboard", http.StatusInternalServerError}
		}

		jsonRes, err := json.Marshal(users)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}
