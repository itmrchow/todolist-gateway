package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/itmrchow/microservice-common/token"
	"github.com/spf13/viper"

	handlers "github.com/itmrchow/todolist-gateway/internal/errors"
)

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		secretKey := viper.GetString("JWT_SECRET_KEY")
		issuer := viper.GetString("JWT_ISSUER")

		userID, err := token.ValidateToken(tokenString, secretKey, issuer)

		if err != nil {
			if errors.Is(err, token.ErrExpiredToken) {
				http.Error(w, handlers.ErrMsg401TokenExpired, http.StatusUnauthorized)
				return
			} else {
				http.Error(w, handlers.ErrMsg401TokenInvalid, http.StatusUnauthorized)
				return
			}
		}

		r.Header.Set("X-User-ID", userID)
		next.ServeHTTP(w, r)
	})
}
