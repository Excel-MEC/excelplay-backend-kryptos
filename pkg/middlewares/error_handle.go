package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
	"github.com/sirupsen/logrus"
)

// ErrorsMiddleware handles the logging and error handling when http errors occur
func ErrorsMiddleware(next httperrors.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if e := next(w, r); e != nil {
			requestLog := fmt.Sprintf("%s\t%s\t%s",
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
	})
}
