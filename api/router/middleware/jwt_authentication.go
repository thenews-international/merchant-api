package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"merchant/model"
)

const (
	jwtErrInvalidToken     = "invalid token"
	jwtErrInvalidTokenUser = "invalid token user"

	emptyUuid = "00000000-0000-0000-0000-000000000000"
)

var (
	secret            = "pacenow-token"
	jwtErrAlgoInvalid = errors.New("JWT algorithm mismatch")
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) < 7 || strings.ToUpper(tokenString[0:6]) != "BEARER" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error": "%v"}`, jwtErrInvalidToken)
			return
		}

		tokenString = tokenString[7:]
		token := &model.Token{}

		_, err := jwt.ParseWithClaims(tokenString, token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwtErrAlgoInvalid
			}

			return []byte(secret), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error": "%v"}`, err.Error())
			return
		}

		if token.UserId == "" || token.UserId == emptyUuid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error": "%v"}`, jwtErrInvalidTokenUser)
			return
		}

		ctx := context.WithValue(r.Context(), model.CtxKeyXUser, token.ToCtxUser())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
