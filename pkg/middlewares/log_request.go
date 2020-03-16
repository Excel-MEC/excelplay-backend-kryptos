package middlewares

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// LoggerMiddleware is a middleware to log all incoming requests
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Println(r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
