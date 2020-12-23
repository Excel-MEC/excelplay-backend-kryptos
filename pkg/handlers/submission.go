package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/liveleaderboard"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
	"github.com/dgrijalva/jwt-go"
)

// Only for swagger documentation, do not use in code.
type swagRequest struct {
	Answer string `json:"answer" example:"excel"`
}

// HandleSubmission handles answer attempts
// @Summary takes a post request with the answer attempt.
// @Description takes a post request with the answer attempt.
// @Tags Kryptos
// @Accept json
// @Produce plain
// @Param payload body swagRequest true "Answer format"
// @Success 200 {object} string "Returns 'success' for correct answer, 'fail' for wrong answer."
// @Failure 500 {string} string
// @Router /api/submit [post]
func HandleSubmission(db *database.DB, env *env.Config) httperrors.Handler {
	type request struct {
		Answer string `json:"answer"`
	}

	// Used to prevent multiple submissions from the same user
	// CAUTION: will not work if backend is scaled to multiple instances. In such cases,
	// create a separate, common service for implementing this functionality.
	pendingTransactions := make(map[int]bool)
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		// Obtain values from JWT
		props, _ := r.Context().Value("props").(jwt.MapClaims)
		userID, _ := strconv.Atoi(props["user_id"].(string))

		// Expected POST format is { "answer": "attempt" }
		input := json.NewDecoder(r.Body)
		input.DisallowUnknownFields()

		var req request
		err := input.Decode(&req)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not deserialize json", http.StatusInternalServerError}
		}

		var currUser database.User
		err = db.GetUser(&currUser, userID)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve user", http.StatusInternalServerError}
		}

		_, err = db.LogAnswerAttempt(userID, currUser, req.Answer)

		var correctAns string
		err = db.GetCorrectAns(currUser, &correctAns)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve the answer", http.StatusInternalServerError}
		}

		// If the user has already submitted the correct answer to their current level,
		// wait until their level has been incremented in the DB before allowing them to attempt again.
		_, userTransactionPending := pendingTransactions[userID]
		if req.Answer == correctAns && !userTransactionPending {
			pendingTransactions[userID] = true
			_, err := db.CorrectAnswerSubmitted(userID)
			delete(pendingTransactions, userID)

			if err != nil {
				return &httperrors.HTTPError{r, err, "Could not update user progress", http.StatusInternalServerError}
			}
			// Send update to leaderboard
			liveleaderboard.UpdateUser <- userID
			jsonRes, err := json.Marshal(map[string]string{"answer": "correct"})
			if err != nil {
				return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write(jsonRes)
			return nil
		}

		jsonRes, err := json.Marshal(map[string]string{"answer": "wrong"})
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(jsonRes)
		return nil
	}
}
