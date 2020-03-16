package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (s *server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin"))
}

// heartBeat sends a response back, just to check if the server is up.
func (s *server) heartBeat() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) *httpError {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test"))
		return nil
	}
}

func (s *server) handleNextQuestion() httpHandler {
	// Values that can be nil or a non-nullable value,
	// such as a string are given the empty interface type
	type response struct {
		Number     int         `json:"number" db:"number"`
		Question   interface{} `json:"question" db:"question"`
		ImageLevel bool        `json:"image_level" db:"image_level"`
		LevelFile  interface{} `json:"level_file" db:"level_file"`
		Hints      []string    `json:"hints"`
	}
	return func(w http.ResponseWriter, r *http.Request) *httpError {
		// replace when auth is ready
		uuid := "c327ea2c-6539-11ea-8c85-0242ac190002"

		var currLev int
		err := s.db.Get(&currLev, "select curr_level from kuser where id = $1", uuid)
		if err != nil {
			return &httpError{r, err, "Could not retrieve curr_level", http.StatusInternalServerError}
		}

		var res response
		// Select all attributes except the answer
		err = s.db.Get(&res, "select number, question, image_level, level_file from levels where number = $1", currLev)
		if err != nil {
			return &httpError{r, err, "Could not retrieve question details", http.StatusInternalServerError}
		}

		var hints []string
		err = s.db.Select(&hints, "select content from hints where number = $1", currLev)
		if err != nil {
			return &httpError{r, err, "Could not retrieve hints", http.StatusInternalServerError}
		}
		res.Hints = hints

		jsonRes, err := json.Marshal(res)
		if err != nil {
			return &httpError{r, err, "Could not serialize json", http.StatusInternalServerError}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}

func (s *server) handleSubmission() httpHandler {
	type request struct {
		Answer string `json:"answer"`
	}
	type user struct {
		Name      string `db:"name"`
		CurrLevel int    `db:"curr_level"`
	}
	return func(w http.ResponseWriter, r *http.Request) *httpError {
		// replace when auth is ready
		uuid := "c327ea2c-6539-11ea-8c85-0242ac190002"

		// Expected POST format is { "answer": "attempt" }
		input := json.NewDecoder(r.Body)
		input.DisallowUnknownFields()

		var req request
		err := input.Decode(&req)
		if err != nil {
			return &httpError{r, err, "Could not deserialize json", http.StatusInternalServerError}
		}

		var currUser user
		err = s.db.Get(&currUser, "select name, curr_level from kuser where id = $1", uuid)
		if err != nil {
			return &httpError{r, err, "Could not retrieve user", http.StatusInternalServerError}
		}

		_, err = s.db.Exec("insert into answer_logs values($1, $2, $3, $4)", uuid, currUser.Name, req.Answer, time.Now())

		var correctAns string
		err = s.db.Get(&correctAns, "select answer from levels where number = $1", currUser.CurrLevel)
		if err != nil {
			return &httpError{r, err, "Could not retrieve the answer", http.StatusInternalServerError}
		}

		if req.Answer == correctAns {
			_, err := s.db.Exec("update kuser set curr_level = curr_level + 1 where id = $1", uuid)
			if err != nil {
				return &httpError{r, err, "Could not update user progress", http.StatusInternalServerError}
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
