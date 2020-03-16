package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func (s *server) authMiddleware(next httpHandler) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) *httpError {
		// TODO: Sample JWT, remove later.
		// jwtToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE1ODQyNzYxNDYsImV4cCI6MTYxNTgxMjE0NSwiYXVkIjoiIiwic3ViIjoiYzMyN2VhMmMtNjUzOS0xMWVhLThjODUtMDI0MmFjMTkwMDAyIiwibmFtZSI6IkpvaG4gRG9lIn0.f94bHZLazkHYNqYMgaHIpPIF4WLFQkfR3rqvN3KiIC9egLI-jf_HJTPbiNLby0SMB1el7im4VS8tG6Uq6p3TWw"
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			return &httpError{r, errors.New("Malformed token"), "Malformed token", http.StatusUnauthorized}
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(SECRETKEY), nil
			})
			if err != nil {
				return &httpError{r, err, "Malformed token", http.StatusUnauthorized}
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				return &httpError{r, errors.New("Failed to validate claims"), "Unauthorized", http.StatusUnauthorized}
			}
		}
		return nil
	}
}
