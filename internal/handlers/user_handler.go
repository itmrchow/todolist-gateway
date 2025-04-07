package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/itmrchow/todolist-proto/protobuf/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/itmrchow/todolist-gateway/internal/dto"
	mErr "github.com/itmrchow/todolist-gateway/internal/errors"
	"github.com/itmrchow/todolist-gateway/utils"
)

var _ UserHandlerInterface = &UserHandler{}

type UserHandler struct {
	validate   *validator.Validate
	userClient user.UserServiceClient
}

func NewUserHandler(validate *validator.Validate, userClient user.UserServiceClient) *UserHandler {
	return &UserHandler{
		validate:   validate,
		userClient: userClient,
	}
}

func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var req dto.RegisterUserReqDTO
	var resp dto.BaseRespDTO

	// request handler
	if err := utils.HandleRequest(r, w, &req, u.validate); err != nil {
		return
	}

	_, err := u.userClient.Register(r.Context(), &user.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.AlreadyExists:
				resp.Message = mErr.ErrMsg409Conflict
				resp.Data = s.Message()
				utils.ResponseWriter(r, w, http.StatusConflict, resp) // 409
			default:
				resp.InternalErrorResp(r, err)
				utils.ResponseWriter(r, w, http.StatusInternalServerError, resp) // 500
			}
		} else {
			resp.InternalErrorResp(r, err)
			utils.ResponseWriter(r, w, http.StatusConflict, resp) // 500

		}
		return
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
