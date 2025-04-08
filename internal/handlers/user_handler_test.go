package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/itmrchow/todolist-proto/protobuf"
	"github.com/itmrchow/todolist-proto/protobuf/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/itmrchow/todolist-gateway/internal/dto"
	mErr "github.com/itmrchow/todolist-gateway/internal/errors"
)

func TestRegisterUser(t *testing.T) {

	validate := validator.New()
	mockUserClient := user.NewMockUserServiceClient(t)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name       string
		args       args
		mockFunc   func()
		assertFunc func(t *testing.T, w http.ResponseWriter)
	}{
		{
			name: "request body is empty",
			args: func() args {

				reqDtoJSON, _ := json.Marshal("")

				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/register", body)
				return args{w: w, r: r}
			}(),
			mockFunc: func() {},
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)

				// status code
				assert.Equal(t, recorder.Code, http.StatusBadRequest)

				// response body
				var respDto dto.BaseRespDTO[string]
				err := json.NewDecoder(recorder.Body).Decode(&respDto)
				assert.NoError(t, err)

				// assert response body
				assert.Equal(t, respDto.Message, mErr.ErrMsg400BadRequest)
				assert.NotNil(t, respDto.Data)
			},
		},
		{
			name: "request body input error",
			args: func() args {
				reqDto := dto.RegisterUserReqDTO{
					Email:    "test@example",
					Name:     "testtesttesttesttesttesttesttest",
					Password: "test1234test1234test1234test1234test1234test1234test1234test1234",
				}

				reqDtoJSON, _ := json.Marshal(reqDto)
				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/register", body)
				return args{w: w, r: r}
			}(),
			mockFunc: func() {},
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)

				// status code
				assert.Equal(t, recorder.Code, http.StatusBadRequest)

				// response body
				var respDto dto.BaseRespDTO[[]dto.FieldError]
				err := json.NewDecoder(recorder.Body).Decode(&respDto)
				assert.NoError(t, err)

				// assert response body
				assert.Equal(t, respDto.Message, mErr.ErrMsg400BadRequest)
				// Assert validation error format
				assert.Equal(t, respDto.Data[0].Key, "Email")
				assert.Equal(t, respDto.Data[0].Error, "email")
				assert.Equal(t, respDto.Data[1].Key, "Name")
				assert.Equal(t, respDto.Data[1].Error, "max")
				assert.Equal(t, respDto.Data[2].Key, "Password")
				assert.Equal(t, respDto.Data[2].Error, "max")
			},
		},
		{
			name: "grpc client error",
			args: func() (a args) {
				// request dto
				reqDto := dto.RegisterUserReqDTO{
					Email:    "test@example.com",
					Name:     "test",
					Password: "test1234",
				}

				reqDtoJSON, _ := json.Marshal(reqDto)
				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/register", body)
				return args{w: w, r: r}

			}(),
			mockFunc: func() {
				mockUserClient.EXPECT().Register(mock.Anything, mock.Anything).Return(nil, errors.New("some error")).Once()
			},
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)

				// status code
				assert.Equal(t, recorder.Code, http.StatusInternalServerError)

				// response body
				var respDto dto.BaseRespDTO[string]
				err := json.NewDecoder(recorder.Body).Decode(&respDto)
				assert.NoError(t, err)

				// assert response body
				assert.Equal(t, respDto.Message, mErr.ErrMsg500InternalServerError)
				assert.Equal(t, respDto.Data, "")
			},
		},
		{
			name: "email already exists",
			args: func() (a args) {
				// request dto
				reqDto := dto.RegisterUserReqDTO{
					Email:    "test@example.com",
					Name:     "test",
					Password: "test1234",
				}

				reqDtoJSON, _ := json.Marshal(reqDto)
				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/register", body)
				return args{w: w, r: r}

			}(),
			mockFunc: func() {
				mockUserClient.EXPECT().Register(mock.Anything, mock.Anything).Return(nil, status.Error(codes.AlreadyExists, "exists msg")).Once()
			},
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)

				// status code
				assert.Equal(t, recorder.Code, http.StatusConflict)

				// response body
				var respDto dto.BaseRespDTO[string]
				err := json.NewDecoder(recorder.Body).Decode(&respDto)
				assert.NoError(t, err)

				// assert response body
				assert.Equal(t, respDto.Message, mErr.ErrMsg409Conflict)
				assert.Equal(t, respDto.Data, "exists msg")
			},
		},
		{
			name: "internal server error",
			args: func() (a args) {
				// request dto
				reqDto := dto.RegisterUserReqDTO{
					Email:    "test@example.com",
					Name:     "test",
					Password: "test1234",
				}

				reqDtoJSON, _ := json.Marshal(reqDto)
				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/register", body)
				return args{w: w, r: r}

			}(),
			mockFunc: func() {
				mockUserClient.EXPECT().Register(mock.Anything, mock.Anything).Return(nil, status.Error(codes.Internal, "some error")).Once()
			},
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)

				// status code
				assert.Equal(t, recorder.Code, http.StatusInternalServerError)

				// response body
				var respDto dto.BaseRespDTO[string]
				err := json.NewDecoder(recorder.Body).Decode(&respDto)
				assert.NoError(t, err)

				// assert response body
				assert.Equal(t, respDto.Message, mErr.ErrMsg500InternalServerError)
				assert.Equal(t, respDto.Data, "")
			},
		},
		{
			name: "register user success",
			args: func() (a args) {
				// request dto
				reqDto := dto.RegisterUserReqDTO{
					Email:    "test@example.com",
					Name:     "test",
					Password: "test1234",
				}

				reqDtoJSON, _ := json.Marshal(reqDto)
				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/register", body)
				return args{w: w, r: r}

			}(),
			mockFunc: func() {
				mockUserClient.EXPECT().Register(mock.Anything, mock.Anything).Return(&protobuf.EmptyResponse{}, nil).Once()
			},
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)

				// status code
				assert.Equal(t, recorder.Code, http.StatusCreated)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				validate:   validate,
				userClient: mockUserClient,
			}
			tt.mockFunc()
			u.RegisterUser(tt.args.w, tt.args.r)
			tt.assertFunc(t, tt.args.w)
		})
	}
}

func TestLoginUser(t *testing.T) {

	validate := validator.New()
	mockUserClient := user.NewMockUserServiceClient(t)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name       string
		args       args
		mockFunc   func()
		assertFunc func(t *testing.T, w http.ResponseWriter)
	}{
		{
			name: "request body input error",
			args: func() (a args) {
				reqDto := dto.LoginUserReqDTO{
					Email:    "test@example",
					Password: "test1234test1234test1234test1234test1234test1234test1234test1234",
				}

				reqDtoJSON, _ := json.Marshal(reqDto)
				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/register", body)
				return args{w: w, r: r}
			}(),
			mockFunc: func() {
			},
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)

				// status code
				assert.Equal(t, recorder.Code, http.StatusBadRequest)

				// response body
				var respDto dto.BaseRespDTO[[]dto.FieldError]
				err := json.NewDecoder(recorder.Body).Decode(&respDto)
				assert.NoError(t, err)

				// assert response body
				assert.Equal(t, respDto.Message, mErr.ErrMsg400BadRequest)
				assert.Equal(t, respDto.Data[0].Key, "Email")
				assert.Equal(t, respDto.Data[0].Error, "email")
				assert.Equal(t, respDto.Data[1].Key, "Password")
				assert.Equal(t, respDto.Data[1].Error, "max")
			},
		},
		{
			name: "svc return unauthenticated",
			args: func() (a args) {
				reqDto := dto.LoginUserReqDTO{
					Email:    "test@example.com",
					Password: "test1234",
				}

				reqDtoJSON, _ := json.Marshal(reqDto)
				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/login", body)
				return args{w: w, r: r}
			}(),
			mockFunc: func() {
				mockUserClient.EXPECT().Login(mock.Anything, mock.Anything).Return(nil, status.Error(codes.Unauthenticated, "unauthenticated")).Once()
			},
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)

				// status code
				assert.Equal(t, recorder.Code, http.StatusUnauthorized)

				// response body
				var respDto dto.BaseRespDTO[string]
				err := json.NewDecoder(recorder.Body).Decode(&respDto)
				assert.NoError(t, err)

				// assert response body
				assert.Equal(t, respDto.Message, mErr.ErrMsg401AuthFailed)
				assert.Equal(t, respDto.Data, "unauthenticated")
			},
		},
		{
			name: "svc return internal server error",
			args: func() (a args) {
				reqDto := dto.LoginUserReqDTO{
					Email:    "test@example.com",
					Password: "test1234",
				}

				reqDtoJSON, _ := json.Marshal(reqDto)
				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/login", body)
				return args{w: w, r: r}
			}(),
			mockFunc: func() {
				mockUserClient.EXPECT().Login(mock.Anything, mock.Anything).Return(nil, status.Error(codes.Internal, "internal error")).Once()
			},
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)

				// status code
				assert.Equal(t, recorder.Code, http.StatusInternalServerError)

				// response body
				var respDto dto.BaseRespDTO[string]
				err := json.NewDecoder(recorder.Body).Decode(&respDto)
				assert.NoError(t, err)

				// assert response body
				assert.Equal(t, respDto.Message, mErr.ErrMsg500InternalServerError)
				assert.Equal(t, respDto.Data, "")
			},
		},
		{
			name: "svc internal server error",
			args: func() (a args) {
				reqDto := dto.LoginUserReqDTO{
					Email:    "test@example.com",
					Password: "test1234",
				}

				reqDtoJSON, _ := json.Marshal(reqDto)
				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/login", body)
				return args{w: w, r: r}
			}(),
			mockFunc: func() {
				mockUserClient.EXPECT().Login(mock.Anything, mock.Anything).Return(nil, errors.New("some error")).Once()
			},
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)

				// status code
				assert.Equal(t, recorder.Code, http.StatusInternalServerError)

				// response body
				var respDto dto.BaseRespDTO[string]
				err := json.NewDecoder(recorder.Body).Decode(&respDto)
				assert.NoError(t, err)

				// assert response body
				assert.Equal(t, respDto.Message, mErr.ErrMsg500InternalServerError)
				assert.Equal(t, respDto.Data, "")
			},
		},
		{
			name: "success",
			args: func() (a args) {
				reqDto := dto.LoginUserReqDTO{
					Email:    "test@example.com",
					Password: "test1234",
				}

				reqDtoJSON, _ := json.Marshal(reqDto)
				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/login", body)
				return args{w: w, r: r}
			}(),
			mockFunc: func() {
				mockUserClient.EXPECT().Login(mock.Anything, mock.Anything).Return(&user.LoginResponse{
					Id:        "123",
					Name:      "test",
					Email:     "test@example.com",
					Token:     "test-token",
					ExpiresIn: &timestamppb.Timestamp{Seconds: 1234567890},
				}, nil).Once()
			},
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)

				// status code
				assert.Equal(t, recorder.Code, http.StatusOK)

				// response body
				var respDto dto.BaseRespDTO[dto.LoginUserRespDTO]
				err := json.NewDecoder(recorder.Body).Decode(&respDto)
				assert.NoError(t, err)

				// assert response body
				assert.Equal(t, respDto.Message, "SUCCESS")
				assert.Equal(t, respDto.Data.ID, "123")
				assert.Equal(t, respDto.Data.Name, "test")
				assert.Equal(t, respDto.Data.Email, "test@example.com")
				assert.Equal(t, respDto.Data.Token, "test-token")
				assert.Equal(t, respDto.Data.ExpiresIn.Unix(), int64(1234567890))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				validate:   validate,
				userClient: mockUserClient,
			}

			tt.mockFunc()
			u.LoginUser(tt.args.w, tt.args.r)
			tt.assertFunc(t, tt.args.w)
		})
	}
}
