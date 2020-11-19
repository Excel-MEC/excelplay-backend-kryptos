package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
	"github.com/dgrijalva/jwt-go"
)

// Only for swagger documentation, do not use in code.
type swagQresponse struct {
	Number     int      `json:"number" example:"1"`
	Question   string   `json:"question" example:"What is MEC's techfest?"`
	ImageLevel bool     `json:"image_level" example:"true"`
	LevelFile  string   `json:"level_file" example:"url_of_image"`
	Hints      []string `json:"hints" example:"['Hint 1', 'Hint 2']"`
}

// HandleNextQuestion handles any request made to the /api/question/ endpoint
// @Summary returns the question for the level the user is on.
// @Description Sends back the question for the level the user is on. If this is a new user, a user instance is created in the DB and the first question is returned.
// @Tags Kryptos
// @Produce json
// @Success 200 {object} swagQresponse "Returns the question and it's details."
// @Failure 500 {string} string
// @Router /api/question [get]
func HandleNextQuestion(db *database.DB, env *env.Config) httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		// Obtain values from JWT
		props, _ := r.Context().Value("props").(jwt.MapClaims)
		name := props["name"].(string)
		userID, _ := strconv.Atoi(props["user_id"].(string))

		var currLev int
		err := db.GetCurrLevel(userID, &currLev)
		if err != nil && err.Error() == "sql: no rows in result set" {
			_, err := db.CreateNewUser(userID, name)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Could not create new user", http.StatusInternalServerError}
			}
			db.GetCurrLevel(userID, &currLev)
		} else if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve curr_level", http.StatusInternalServerError}
		}

		var res database.QResponse
		err = db.GetQuestion(currLev, &res)
		if err != nil {
			fmt.Println(err.Error())
			return &httperrors.HTTPError{r, err, "Could not retrieve question details", http.StatusInternalServerError}
		}

		var hints []string
		err = db.GetHints(currLev, &hints)
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
