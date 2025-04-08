package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/itmrchow/microservice-common/token"
	"github.com/spf13/viper"

	"github.com/itmrchow/todolist-gateway/internal/dto"
	mErr "github.com/itmrchow/todolist-gateway/internal/errors"
	"github.com/itmrchow/todolist-gateway/utils"
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
				resp := dto.NewErrorResponse(mErr.ErrMsg401TokenExpired, err.Error())
				utils.ResponseWriter(r, w, http.StatusUnauthorized, resp)
			} else {
				resp := dto.NewErrorResponse(mErr.ErrMsg401TokenInvalid, err.Error())
				utils.ResponseWriter(r, w, http.StatusUnauthorized, resp)
			}

			return
		}

		r.Header.Set("X-User-ID", userID)
		next.ServeHTTP(w, r)
	})
}
