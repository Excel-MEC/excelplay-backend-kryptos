package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
)

// Only for swagger documentation, do not use in code.
type swagUser struct {
	Name      string `json:"name" example:"Aswin G"`
	CurrLevel int    `json:"curr_level" example:"18"`
	ProPic    string `json:"profile_pic" example:"https://youtu.be/dQw4w9WgXcQ"`
}

// GetLeaderboard returns all the users ordered in descending order of level,
// and for users on the same level, in the ascending order of last submission time.
// @Summary return leaderboard.
// @Description Sends back the leaderboard in descending order of level, and for users on the same level, in the ascending order of last successful submission time.
// @Tags Kryptos
// @Produce json
// @Success 200 {object} []swagUser "Returns the leaderboard"
// @Failure 500 {string} string "Could not serialize json"
// @Router /api/leaderboard [get]
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
