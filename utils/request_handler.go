package utils

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/itmrchow/todolist-gateway/internal/dto"
	mErr "github.com/itmrchow/todolist-gateway/internal/errors"
)

func HandleRequest(r *http.Request, w http.ResponseWriter, req interface{}, validate *validator.Validate) error {

	// 解析請求體
	if err := DecodeJSONBody(r, req); err != nil {

		resp := dto.NewErrorResponse(mErr.ErrMsg400BadRequest, err.Error())
		ResponseWriter(r, w, http.StatusBadRequest, resp)

		return err
	}

	// 驗證請求
	if err := validate.Struct(req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			resp := dto.NewFieldErrorRespDTO(err.(validator.ValidationErrors))
			ResponseWriter(r, w, http.StatusBadRequest, resp)
		} else {
			resp := dto.NewInternalErrorRespDTO(r.Header.Get("X-Trace-ID"), err)
			ResponseWriter(r, w, http.StatusInternalServerError, resp)
		}

		return err
	}

	return nil
}
