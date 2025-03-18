package handlers

import "github.com/gorilla/mux"

func RegisterHandlers() (router *mux.Router) {
	router = mux.NewRouter()

	// middleware

	// router
	RegisterUserRouter(router)
	RegisterTaskRouter(router)

	return
}

func RegisterUserRouter(r *mux.Router) {
	r.HandleFunc("/users/register", RegisterUser).Methods("POST")
	// r.HandleFunc("/users/login", LoginUser).Methods("POST")
}

func RegisterTaskRouter(r *mux.Router) {
	// r.HandleFunc("/tasks/create", CreateTask).Methods("POST")
	// r.HandleFunc("/tasks/update/{id}", UpdateTask).Methods("PUT")
	// r.HandleFunc("/tasks/delete/{id}", DeleteTask).Methods("DELETE")
	// r.HandleFunc("/tasks/list", ListTasks).Methods("GET")
}
