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

	err := utils.DecodeJSONBody(r, &req)
	if err != nil {
		resp.Message = mErr.ErrMsg400BadRequest
		resp.Data = err.Error()

		utils.ResponseWriter(r, w, http.StatusBadRequest, resp)
		return
	}

	err = u.validate.Struct(req)
	if err != nil {
		// 400
		resp.Message = mErr.ErrMsg400BadRequest

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors []map[string]string
			for _, fieldError := range validationErrors {
				errors = append(errors, map[string]string{
					"key":   fieldError.Field(),
					"error": fieldError.Tag(),
				})
			}
			resp.Data = errors // 將所有驗證錯誤信息放入
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
	panic("TODO: Implement")
}
