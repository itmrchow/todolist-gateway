package handlers

import "net/http"

type UserHandlerInterface interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}
