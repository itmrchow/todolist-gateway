package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/itmrchow/todolist-system/utils"
)

var _ UserHandlerInterface = &UserHandler{}

type UserHandler struct {
	validate *validator.Validate
}

func NewUserHandler(validate *validator.Validate) *UserHandler {
	return &UserHandler{
		validate: validate,
	}
}

func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterUserReqDTO
	err := utils.DecodeJSONBody(r, &req)
	if err != nil {
		// 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.validate.Struct(req)
	if err != nil {
		// 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 201
}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	panic("TODO: Implement")
}
