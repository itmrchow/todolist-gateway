package dto

import (
	"github.com/go-playground/validator/v10"

	mErr "github.com/itmrchow/todolist-gateway/internal/errors"
)

type BaseRespDTO struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (resp *BaseRespDTO) ValidatorErrorResp(err validator.ValidationErrors) {
	var errors []map[string]string
	for _, fieldError := range err {
		errors = append(errors, map[string]string{
			"key":   fieldError.Field(),
			"error": fieldError.Tag(),
		})
	}
	resp.Message = mErr.ErrMsg400BadRequest
	resp.Data = errors // 將所有驗證錯誤信息放入
}

// type FailedRespDTO struct {
// 	BaseRespDTO
// 	Error string `json:"error"`
// }

// type SuccessRespDTO struct {
// 	BaseRespDTO
// }
