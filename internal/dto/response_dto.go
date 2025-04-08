package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	mErr "github.com/itmrchow/todolist-gateway/internal/errors"
)

type BaseRespDTO[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type FieldError struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}

func NewFieldErrorRespDTO(err validator.ValidationErrors) (respDTO BaseRespDTO[[]FieldError]) {
	respDTO.Message = mErr.ErrMsg400BadRequest

	var errors []FieldError
	for _, fieldError := range err {
		errors = append(errors, FieldError{
			Key:   fieldError.Field(),
			Error: fieldError.Tag(),
		})
	}

	respDTO.Data = errors
	return
}

func NewInternalErrorRespDTO(traceID string, err error) (respDTO BaseRespDTO[any]) {
	log.Error().Err(err).
		Str("trace_id", traceID).
		Msg("service internal error")

	respDTO.Message = mErr.ErrMsg500InternalServerError
	return
}

func NewErrorResponse(message string, err string) (respDTO BaseRespDTO[string]) {

	respDTO.Message = message
	respDTO.Data = err

	return
}

func NewSuccessResponse[T any](data T) (respDTO BaseRespDTO[T]) {
	respDTO.Message = "SUCCESS"
	respDTO.Data = data
	return
}
