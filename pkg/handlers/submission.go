package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
)

// HandleSubmission handles any answer submission made on the /api/submit/ endpoint
func HandleSubmission(db *database.DB, env *env.Config) httperrors.Handler {
	type request struct {
		Answer string `json:"answer"`
	}
	type user struct {
		Name      string `db:"name"`
		CurrLevel int    `db:"curr_level"`
	}
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		// replace when auth is ready
		uuid := "c327ea2c-6539-11ea-8c85-0242ac190002"

		// Expected POST format is { "answer": "attempt" }
		input := json.NewDecoder(r.Body)
		input.DisallowUnknownFields()

		var req request
		err := input.Decode(&req)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not deserialize json", http.StatusInternalServerError}
		}

		var currUser user
		err = db.Get(&currUser, "select name, curr_level from kuser where id = $1", uuid)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve user", http.StatusInternalServerError}
		}

		_, err = db.Exec("insert into answer_logs values($1, $2, $3, $4)", uuid, currUser.Name, req.Answer, time.Now())

		var correctAns string
		err = db.Get(&correctAns, "select answer from levels where number = $1", currUser.CurrLevel)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve the answer", http.StatusInternalServerError}
		}

		if req.Answer == correctAns {
			_, err := db.Exec("update kuser set curr_level = curr_level + 1 where id = $1", uuid)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Could not update user progress", http.StatusInternalServerError}
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))
			return nil
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("fail"))
		return nil
	}
}
