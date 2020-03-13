package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin"))
}

// heartBeat sends a response back, just to check if the server is up.
func (s *server) heartBeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Test"))
}

func (s *server) handleNextQuestion() http.HandlerFunc {
	// Values that can be nil or a non-nullable value,
	// such as a string are given the empty interface type
	type response struct {
		Number     int         `json:"number" db:"number"`
		Question   interface{} `json:"question" db:"question"`
		ImageLevel bool        `json:"image_level" db:"image_level"`
		LevelFile  interface{} `json:"level_file" db:"level_file"`
		Hints      []string    `json:"hints"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		requestLog := fmt.Sprintf("%s\t%s",
			r.Method,
			r.RequestURI,
		)
		logger := s.logger.WithField("request", requestLog)

		// replace when auth is ready
		uuid := "c327ea2c-6539-11ea-8c85-0242ac190002"

		var currLev int
		err := s.db.Get(&currLev, "select curr_level from kuser where id = $1", uuid)
		if err != nil {
			logger.Errorf(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var res response
		// Select all attributes except the answer
		err = s.db.Get(&res, "select number, question, image_level, level_file from levels where number = $1", currLev)
		if err != nil {
			logger.Errorf(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var hints []string
		err = s.db.Select(&hints, "select content from hints where number = $1", currLev)
		if err != nil {
			logger.Errorf(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Hints = hints

		jsonRes, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return
	}
}

func (s *server) handleSubmission() http.HandlerFunc {
	type response struct {
		Question   string   `json:"question"`
		ImageLevel bool     `json:"image_level"`
		LevelFile  string   `json:"level_file"`
		Hints      []string `json:"hints"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// do handling
	}
}
