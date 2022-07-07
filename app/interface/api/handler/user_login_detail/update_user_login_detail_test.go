package user_login_detail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	uldMock "www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/mock"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

func TestUserLoginDetailHandler_UpdateUserLoginDetail(t *testing.T) {
	log.MakeLogger("tmp/infotax.log", true)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var ctx gin.Context

	validDetail := in.UpdateUserDetail{
		DomainName: "primzz",
		EmailID:    "primzz@gmail.com",
		Role:       "Payroll_User",
	}

	invalidDetail := in.UpdateUserDetail{
		DomainName: "",
		EmailID:    "primzz",
		Role:       "primzz",
	}

	dtl := in.UpdateUserDetail{
		DomainName: "primzz",
		EmailID:    "primzz@gmail.com",
		Role:       "Payroll_User",
	}

	id := "2521"

	t.Run("Failure, While request payload have invalid content type", func(t *testing.T) {
		uc := uldMock.NewMockUseCaser(ctrl)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.PUT("/infotax/user_login_detail/:emp_id", h.UpdateUserLoginDetail)

		w := httptest.NewRecorder()
		reqBody, err := json.Marshal(validDetail)
		assert.NoError(t, err)
		r, err := http.NewRequest("PUT", "/infotax/user_login_detail/"+id, bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusUnsupportedMediaType, result.StatusCode) {
			msg := fmt.Sprint("/infotax/user_login_detail/:emp_id only accepts Content-Type application/json")
			respBody, err := json.Marshal(gin.H{
				"error": error.NewUnsupportMediaType(msg),
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Failure, While request payload have invalid values", func(t *testing.T) {
		uc := uldMock.NewMockUseCaser(ctrl)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.PUT("/infotax/user_login_detail/:emp_id", h.UpdateUserLoginDetail)

		w := httptest.NewRecorder()
		reqBody, err := json.Marshal(invalidDetail)
		assert.NoError(t, err)
		r, err := http.NewRequest("PUT", "/infotax/user_login_detail/"+id, bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		r.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusBadRequest, result.StatusCode) {
			respBody, err := json.Marshal(gin.H{
				"error": error.NewBadRequest("Invalid request parameters. Verify request parameters once."),
				"invalidArgs": []handler.InvalidArgument{
					handler.InvalidArgument{
						Field: "DomainName",
						Value: "",
						Tag:   "required",
						Param: "",
					},
					handler.InvalidArgument{
						Field: "EmailID",
						Value: "primzz",
						Tag:   "email",
						Param: "",
					},
					handler.InvalidArgument{
						Field: "Role",
						Value: "primzz",
						Tag:   "oneof",
						Param: "Payroll_User Accounts_Team Payroll_Team",
					},
				},
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Failure, While perform upload operation with invalid employee id", func(t *testing.T) {
		var expected out.UpdateResponse
		expectedError := error.NewNotFound("employee_id", id)
		uc := uldMock.NewMockUseCaser(ctrl)
		uc.EXPECT().UpdateUserLoginDetail(gomock.AssignableToTypeOf(&ctx), id, dtl).Return(expected, expectedError)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.PUT("/infotax/user_login_detail/:emp_id", h.UpdateUserLoginDetail)

		w := httptest.NewRecorder()
		reqBody, err := json.Marshal(validDetail)
		assert.NoError(t, err)
		r, err := http.NewRequest("PUT", "/infotax/user_login_detail/"+id, bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		r.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusNotFound, result.StatusCode) {
			respBody, err := json.Marshal(gin.H{
				"error": expectedError,
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Success, While new user login details insertion was completed successfully", func(t *testing.T) {
		expected := out.UpdateResponse{
			IsUpdated: "true",
		}
		uc := uldMock.NewMockUseCaser(ctrl)
		uc.EXPECT().UpdateUserLoginDetail(gomock.AssignableToTypeOf(&ctx), id, dtl).Return(expected, nil)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.PUT("/infotax/user_login_detail/:emp_id", h.UpdateUserLoginDetail)

		w := httptest.NewRecorder()
		reqBody, err := json.Marshal(validDetail)
		assert.NoError(t, err)
		r, err := http.NewRequest("PUT", "/infotax/user_login_detail/"+id, bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		r.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusOK, result.StatusCode) {
			respBody, err := json.Marshal(expected)
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})
}
