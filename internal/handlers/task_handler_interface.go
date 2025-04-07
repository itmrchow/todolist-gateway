package handlers

import "net/http"

type TaskHandlerInterface interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	ListTasks(w http.ResponseWriter, r *http.Request)
}
