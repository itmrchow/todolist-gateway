package handlers

import (
	"github.com/gorilla/mux"
)

func RegisterUserRouter(r *mux.Router, userHandler *UserHandler) {
	r.HandleFunc("/users/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", userHandler.LoginUser).Methods("POST")
}

func RegisterTaskRouter(r *mux.Router, taskHandler *TaskHandler) {
	r.HandleFunc("/tasks/create", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/update/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/delete/{id}", taskHandler.DeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/list", taskHandler.ListTasks).Methods("GET")
}
