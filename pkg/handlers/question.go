package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
	"github.com/dgrijalva/jwt-go"
)

// HandleNextQuestion handles any request made to the /api/question/ endpoint
func HandleNextQuestion(db *database.DB, env *env.Config) httperrors.Handler {
	// Values that can be nil or a non-nullable value,
	// such as a string are given the empty interface type
	type response struct {
		Number     int         `json:"number" db:"number"`
		Question   interface{} `json:"question" db:"question"`
		ImageLevel bool        `json:"image_level" db:"image_level"`
		LevelFile  interface{} `json:"level_file" db:"level_file"`
		Hints      []string    `json:"hints"`
	}
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		// Obtain values from JWT
		props, _ := r.Context().Value("props").(jwt.MapClaims)

		var currLev int
		err := db.Get(&currLev, "select curr_level from kuser where id = $1", props["sub"])
		if err != nil && err.Error() == "sql: no rows in result set" {
			db.Exec("insert into kuser values($1,$2,$3)", props["sub"], props["name"], 1)
		} else if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve curr_level", http.StatusInternalServerError}
		}

		var res response
		// Select all attributes except the answer
		err = db.Get(&res, "select number, question, image_level, level_file from levels where number = $1", currLev)
		if err != nil {
			fmt.Println(err.Error())
			return &httperrors.HTTPError{r, err, "Could not retrieve question details", http.StatusInternalServerError}
		}

		var hints []string
		err = db.Select(&hints, "select content from hints where number = $1", currLev)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve hints", http.StatusInternalServerError}
		}
		res.Hints = hints

		jsonRes, err := json.Marshal(res)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}
