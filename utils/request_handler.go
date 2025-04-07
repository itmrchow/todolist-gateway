package utils

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	"github.com/itmrchow/todolist-gateway/internal/dto"
	mErr "github.com/itmrchow/todolist-gateway/internal/errors"
)

func HandleRequest(r *http.Request, w http.ResponseWriter, req interface{}, validate *validator.Validate) error {
	var resp dto.BaseRespDTO

	// 解析請求體
	if err := DecodeJSONBody(r, req); err != nil {
		resp.Message = mErr.ErrMsg400BadRequest
		resp.Data = err.Error()
		ResponseWriter(r, w, http.StatusBadRequest, resp)
		return err
	}

	// 驗證請求
	if err := validate.Struct(req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			resp.ValidatorErrorResp(err.(validator.ValidationErrors))
		} else {
			resp.Message = mErr.ErrMsg500InternalServerError
			log.Error().Err(err).
				Str("trace_id", r.Header.Get("X-Trace-ID")).
				Msg("request validation error")
		}
		ResponseWriter(r, w, http.StatusBadRequest, resp)
		return err
	}

	return nil
}
