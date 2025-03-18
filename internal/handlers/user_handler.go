package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/itmrchow/todolist-system/utils"
)

var validate *validator.Validate

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterUserReqDTO
	err := utils.DecodeJSONBody(r, &req)
	if err != nil {
		// 400
	}

	err = validate.Struct(req)
	if err != nil {
		// 400
	}

}
