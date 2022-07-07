package user_login_detail

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	uldMock "www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/mock"
)

func TestUserLoginDetailHandler_GetAllUserLoginDetail(t *testing.T) {
	log.MakeLogger("tmp/infotax.log", true)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var ctx gin.Context

	t.Run("Failure, While db error occurred", func(t *testing.T) {
		var dtls []out.UserLoginDetail
		expectedErr := error.NewInternal()
		uc := uldMock.NewMockUseCaser(ctrl)
		uc.EXPECT().GetAllUserLoginDetail(gomock.AssignableToTypeOf(&ctx)).Return(dtls, expectedErr)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.GET("/infotax/user_login_detail/", h.GetAllUserLoginDetail)

		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, "/infotax/user_login_detail/", nil)
		assert.NoError(t, err)
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusInternalServerError, result.StatusCode) {
			respBody, err := json.Marshal(gin.H{
				"error": expectedErr,
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Success, While got all user login details", func(t *testing.T) {
		expected := []out.UserLoginDetail{
			out.UserLoginDetail{
				EmployeeID: "1",
				DomainName: "a",
				EmailID:    "a.s@ivtlinfoview.co.jp",
				Password:   "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
				UUID:       "Testdata",
				IsSignedUp: 0,
				Role:       "Payroll_User",
			},
			out.UserLoginDetail{
				EmployeeID: "b",
				DomainName: "saravase",
				EmailID:    "b.s@ivtlinfoview.co.jp",
				Password:   "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
				UUID:       "Testdata",
				IsSignedUp: 0,
				Role:       "Payroll_User",
			},
		}
		uc := uldMock.NewMockUseCaser(ctrl)
		uc.EXPECT().GetAllUserLoginDetail(gomock.AssignableToTypeOf(&ctx)).Return(expected, nil)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.GET("/infotax/user_login_detail/", h.GetAllUserLoginDetail)

		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, "/infotax/user_login_detail/", nil)
		assert.NoError(t, err)
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
