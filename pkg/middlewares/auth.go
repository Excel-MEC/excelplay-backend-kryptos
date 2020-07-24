package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/strconst"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
	"github.com/dgrijalva/jwt-go"
)

// AuthMiddleware takes a httperrors.Handler and calls it if the JWT in the headers is valid
func AuthMiddleware(next httperrors.Handler, config *env.Config) httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		// TODO: Sample JWT, remove later.
		// jwtToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE1ODQyNzYxNDYsImV4cCI6MTYxNTgxMjE0NSwiYXVkIjoiIiwic3ViIjoiYzMyN2VhMmMtNjUzOS0xMWVhLThjODUtMDI0MmFjMTkwMDAyIiwibmFtZSI6IkpvaG4gRG9lIn0.f94bHZLazkHYNqYMgaHIpPIF4WLFQkfR3rqvN3KiIC9egLI-jf_HJTPbiNLby0SMB1el7im4VS8tG6Uq6p3TWw"
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			return &httperrors.HTTPError{r, errors.New(strconst.MalformedToken), strconst.MalformedToken, http.StatusUnauthorized}
		}
		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Secretkey), nil
		})
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.MalformedToken, http.StatusUnauthorized}
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "props", claims)
			// Access context values in handlers like this
			// props, _ := r.Context().Value("props").(jwt.MapClaims)
			if err := next(w, r.WithContext(ctx)); err != nil {
				return err
			}
		} else {
			return &httperrors.HTTPError{r, errors.New(strconst.ClaimFail), strconst.Unauthorized, http.StatusUnauthorized}
		}
		return nil
	}
}
