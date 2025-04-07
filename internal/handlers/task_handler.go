package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type TaskHandler struct {
	validate *validator.Validate
}

func NewTaskHandler(validate *validator.Validate) *TaskHandler {
	return &TaskHandler{
		validate: validate,
	}
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	panic("TODO: Implement")
}

func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	panic("TODO: Implement")
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	panic("TODO: Implement")
}

func (t *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	panic("TODO: Implement")
}
