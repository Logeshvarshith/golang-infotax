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

func TestUserLoginDetailHandler_CreateUserLoginDetail(t *testing.T) {
	log.MakeLogger("tmp/infotax.log", true)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var ctx gin.Context

	validDetail := in.CreateUserDetail{
		EmployeeID: "222222",
		DomainName: "primzz",
		EmailID:    "primzz@gmail.com",
		Role:       "Payroll_User",
	}

	invalidDetail := in.CreateUserDetail{
		EmployeeID: "",
		DomainName: "",
		EmailID:    "",
		Role:       "",
	}

	dtl := in.CreateUserDetail{
		EmployeeID:   "222222",
		DomainName:   "primzz",
		EmailID:      "primzz@gmail.com",
		Role:         "Payroll_User",
		Password:     "",
		IsSignedUp:   0,
		UUID:         "TestData",
		EnableAccess: "No",
	}

	t.Run("Failure, While request payload have invalid content type", func(t *testing.T) {
		uc := uldMock.NewMockUseCaser(ctrl)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.POST("/infotax/user_login_detail/", h.CreateUserLoginDetail)

		w := httptest.NewRecorder()
		reqBody, err := json.Marshal(validDetail)
		assert.NoError(t, err)
		r, err := http.NewRequest("POST", "/infotax/user_login_detail/", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusUnsupportedMediaType, result.StatusCode) {
			msg := fmt.Sprint("/infotax/user_login_detail/ only accepts Content-Type application/json")
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
		router.POST("/infotax/user_login_detail/", h.CreateUserLoginDetail)

		w := httptest.NewRecorder()
		reqBody, err := json.Marshal(invalidDetail)
		assert.NoError(t, err)
		r, err := http.NewRequest("POST", "/infotax/user_login_detail/", bytes.NewBuffer(reqBody))
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
						Field: "EmployeeID",
						Value: "",
						Tag:   "required",
						Param: "",
					},
					handler.InvalidArgument{
						Field: "DomainName",
						Value: "",
						Tag:   "required",
						Param: "",
					},
					handler.InvalidArgument{
						Field: "EmailID",
						Value: "",
						Tag:   "required",
						Param: "",
					},
					handler.InvalidArgument{
						Field: "Role",
						Value: "",
						Tag:   "required",
						Param: "",
					},
				},
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Failure, While insert available employee id again", func(t *testing.T) {
		var expected out.SaveResponse
		uc := uldMock.NewMockUseCaser(ctrl)
		uc.EXPECT().CreateUserLoginDetail(gomock.AssignableToTypeOf(&ctx), dtl).Return(expected, error.NewConflict("employee_id", dtl.EmployeeID))
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.POST("/infotax/user_login_detail/", h.CreateUserLoginDetail)

		w := httptest.NewRecorder()
		reqBody, err := json.Marshal(validDetail)
		assert.NoError(t, err)
		r, err := http.NewRequest("POST", "/infotax/user_login_detail/", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		r.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusConflict, result.StatusCode) {
			respBody, err := json.Marshal(gin.H{
				"error": error.NewConflict("employee_id", dtl.EmployeeID),
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Success, While new user login details insertion was completed successfully", func(t *testing.T) {
		expected := out.SaveResponse{
			IsSaved: "true",
		}
		uc := uldMock.NewMockUseCaser(ctrl)
		uc.EXPECT().CreateUserLoginDetail(gomock.AssignableToTypeOf(&ctx), dtl).Return(expected, nil)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.POST("/infotax/user_login_detail/", h.CreateUserLoginDetail)

		w := httptest.NewRecorder()
		reqBody, err := json.Marshal(validDetail)
		assert.NoError(t, err)
		r, err := http.NewRequest("POST", "/infotax/user_login_detail/", bytes.NewBuffer(reqBody))
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
