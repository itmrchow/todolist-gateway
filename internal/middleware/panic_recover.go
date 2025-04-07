package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/rs/zerolog/log"
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
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
