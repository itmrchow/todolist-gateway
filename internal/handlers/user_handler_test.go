package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_RegisterUser(t *testing.T) {

	validate := validator.New()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name       string
		args       args
		assertFunc func(t *testing.T, w http.ResponseWriter)
	}{
		{
			name: "request body is empty",
			args: func() args {

				reqDtoJSON, _ := json.Marshal("")

				// reqDto := RegisterUserReqDTO{
				// 	Email:    "test@example.com",
				// 	Name:     "test",
				// 	Password: "test1234",
				// }

				body := bytes.NewReader(reqDtoJSON)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/users/register", body)
				return args{w: w, r: r}
			}(),
			assertFunc: func(t *testing.T, resp http.ResponseWriter) {
				recorder := resp.(*httptest.ResponseRecorder)
				assert.Equal(t, recorder.Code, http.StatusBadRequest)

				t.Log("Response Body:", recorder.Body.String())

				// recorder := w.(*httptest.ResponseRecorder)

				// if w.Code != http.StatusCreated {
				// 	t.Errorf("expected status code %d, got %d", http.StatusCreated, w.Code)
				// }
			},
		},
		// {
		// 	name: "register user success",
		// 	args: func() args {
		// 		reqDto := RegisterUserReqDTO{
		// 			Email:    "test@example.com",
		// 			Name:     "test",
		// 			Password: "test1234",
		// 		}

		// 		reqDtoJSON, _ := json.Marshal(reqDto)
		// 		body := bytes.NewReader(reqDtoJSON)

		// 		w := httptest.NewRecorder()
		// 		r := httptest.NewRequest(http.MethodPost, "/users/register", body)
		// 		return args{w: w, r: r}
		// 	}(),
		// 	assertFunc: func(t *testing.T, resp http.ResponseWriter) {
		// 		recorder := resp.(*httptest.ResponseRecorder)
		// 		assert.Equal(t, recorder.Code, http.StatusCreated)

		// 		log.Println(recorder.Body)
		// 		// recorder := w.(*httptest.ResponseRecorder)

		// 		// if w.Code != http.StatusCreated {
		// 		// 	t.Errorf("expected status code %d, got %d", http.StatusCreated, w.Code)
		// 		// }
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				validate: validate,
			}
			u.RegisterUser(tt.args.w, tt.args.r)
			tt.assertFunc(t, tt.args.w)
		})
	}
}
