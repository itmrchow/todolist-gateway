package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/rs/zerolog/log"

	"github.com/itmrchow/todolist-gateway/internal/dto"
	mErr "github.com/itmrchow/todolist-gateway/internal/errors"
	"github.com/itmrchow/todolist-gateway/utils"
)

func PanicRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// log
				stackTrace := string(debug.Stack())
				log.Error().
					Str("path", r.URL.Path).
					Interface("error", err).
					Str("stack_trace", stackTrace).
					Msg("panic")

				// response
				resp := dto.NewErrorResponse(mErr.ErrMsg500InternalServerError, err.(error).Error())
				utils.ResponseWriter(r, w, http.StatusInternalServerError, resp)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
