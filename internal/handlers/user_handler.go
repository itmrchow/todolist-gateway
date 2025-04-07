package handlers

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/itmrchow/todolist-proto/protobuf/user"
	"github.com/rs/zerolog/log"

	"github.com/itmrchow/todolist-gateway/internal/dto"
	mErr "github.com/itmrchow/todolist-gateway/internal/errors"
	"github.com/itmrchow/todolist-gateway/internal/service"
	"github.com/itmrchow/todolist-gateway/utils"
)

var _ UserHandlerInterface = &UserHandler{}

type UserHandler struct {
	validate *validator.Validate
	userSvc  *service.UserService
}

func NewUserHandler(validate *validator.Validate, userSvc *service.UserService) *UserHandler {
	return &UserHandler{
		validate: validate,
		userSvc:  userSvc,
	}
}

func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var req dto.RegisterUserReqDTO
	var resp dto.BaseRespDTO

	err := utils.DecodeJSONBody(r, &req)
	if err != nil {
		resp.Message = mErr.ErrMsg400BadRequest
		resp.Data = err.Error()

		utils.ResponseWriter(r, w, http.StatusBadRequest, resp)
		return
	}

	err = u.validate.Struct(req)
	if err != nil {

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			// 400
			resp.ValidatorErrorResp(err.(validator.ValidationErrors))
		} else {
			// 500
			resp.Message = mErr.ErrMsg500InternalServerError
			log.Error().Err(err).Str("trace_id", r.Header.Get("X-Trace-ID")).Msg("register user error")
		}

		utils.ResponseWriter(r, w, http.StatusBadRequest, resp)
		return
	}

	// call rpc service
	client, err := u.userSvc.NewClient()
	if err != nil {
		resp.Message = mErr.ErrMsg500InternalServerError
		resp.Data = err.Error()

		utils.ResponseWriter(r, w, http.StatusInternalServerError, resp)
		return
	}

	_, err = client.Register(r.Context(), &user.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
		resp.Message = mErr.ErrMsg500InternalServerError
		resp.Data = err.Error()

		utils.ResponseWriter(r, w, http.StatusInternalServerError, resp)
	}

	// 201
	w.WriteHeader(http.StatusCreated)
}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	var req dto.LoginUserReqDTO
	var resp dto.BaseRespDTO

	err := utils.DecodeJSONBody(r, &req)
	if err != nil {
		resp.Message = mErr.ErrMsg400BadRequest
		resp.Data = err.Error()

		utils.ResponseWriter(r, w, http.StatusBadRequest, resp)
		return
	}

	err = u.validate.Struct(req)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			// 400
			resp.ValidatorErrorResp(err.(validator.ValidationErrors))
		} else {
			// 500
			resp.Message = mErr.ErrMsg500InternalServerError
			log.Error().Err(err).Str("trace_id", r.Header.Get("X-Trace-ID")).Msg("register user error")
		}

		utils.ResponseWriter(r, w, http.StatusBadRequest, resp)
		return
	}

	panic("TODO: Implement")
}
