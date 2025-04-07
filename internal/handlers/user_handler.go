package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/itmrchow/todolist-proto/protobuf/user"

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

	// request handler
	if err := utils.HandleRequest(r, w, &req, u.validate); err != nil {
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
	// var resp dto.BaseRespDTO

	// request handler
	if err := utils.HandleRequest(r, w, &req, u.validate); err != nil {
		return
	}

	panic("TODO: Implement")
}
