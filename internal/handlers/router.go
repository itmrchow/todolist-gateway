package handlers

import (
	"github.com/gorilla/mux"

	"github.com/itmrchow/todolist-gateway/internal/middleware"
)

func RegisterUserRouter(r *mux.Router, userHandler *UserHandler) {

	userRouter := r.NewRoute().Subrouter()

	userRouter.HandleFunc("/users/login", userHandler.LoginUser).Methods("POST")
	userRouter.HandleFunc("/users/register", userHandler.RegisterUser).Methods("POST")
}

func RegisterTaskRouter(r *mux.Router, taskHandler *TaskHandler) {

	taskRouter := r.NewRoute().Subrouter()
	taskRouter.Use(middleware.ValidateToken)

	taskRouter.HandleFunc("/tasks/create", taskHandler.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/tasks/update/{id}", taskHandler.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/tasks/delete/{id}", taskHandler.DeleteTask).Methods("DELETE")
	taskRouter.HandleFunc("/tasks/list", taskHandler.ListTasks).Methods("GET")
}
