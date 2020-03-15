package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func (s *server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Sample JWT, remove later.
		// jwtToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE1ODQyNzYxNDYsImV4cCI6MTYxNTgxMjE0NSwiYXVkIjoiIiwic3ViIjoiYzMyN2VhMmMtNjUzOS0xMWVhLThjODUtMDI0MmFjMTkwMDAyIiwibmFtZSI6IkpvaG4gRG9lIn0.f94bHZLazkHYNqYMgaHIpPIF4WLFQkfR3rqvN3KiIC9egLI-jf_HJTPbiNLby0SMB1el7im4VS8tG6Uq6p3TWw"
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(SECRETKEY), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println(claims["sub"], claims["name"])
				next.ServeHTTP(w, r)
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}
