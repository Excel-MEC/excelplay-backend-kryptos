package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type httpError struct {
	Request *http.Request `json:"-"`
	Error   error         `json:"-"`
	Message string        `json:"error"`
	Code    int           `json:"status"`
}

type httpHandler func(http.ResponseWriter, *http.Request) *httpError

func (fn httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		requestLog := fmt.Sprintf("%s/t%s/t%s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
		)
		logrus.WithField("request", requestLog).Errorf(e.Error.Error())

		res, _ := json.Marshal(e)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(e.Code)
		w.Write(res)
	}
}
