package user_login_detail

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	uldMock "www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/mock"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

func TestUserLoginDetailHandler_SearchUserLoginDetail(t *testing.T) {
	log.MakeLogger("tmp/infotax.log", true)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var ctx gin.Context
	invalidQuery := "?employee_id=2521&domain_name=saravase&isSignedup="
	validQuery := "?employee_id=2521&domain_name=saravase&email_id=saravanakumar.s@ivtlinfoview.co.jp&password=$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO&uuid=Testdata&isSignedup=0&role=Payroll_User"
	filterMap := map[string]interface{}{
		"employee_id": "2521",
		"domain_name": "saravase",
		"email_id":    "saravanakumar.s@ivtlinfoview.co.jp",
		"password":    "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
		"uuid":        "Testdata",
		"isSignedup":  0,
		"role":        "Payroll_User",
	}
	dtls := []out.UserLoginDetail{
		out.UserLoginDetail{
			EmployeeID: "2521",
			DomainName: "saravase",
			EmailID:    "saravase.s@ivtlinfoview.co.jp",
			Password:   "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
			UUID:       "Testdata",
			IsSignedUp: 0,
			Role:       "Payroll_User",
		},
	}
	t.Run("Failure, While give invalid request payload values data type", func(t *testing.T) {
		uc := uldMock.NewMockUseCaser(ctrl)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.GET("/infotax/user_login_detail/search", h.SearchUserLoginDetail)

		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "/infotax/user_login_detail/search"+invalidQuery, nil)
		assert.NoError(t, err)
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusBadRequest, result.StatusCode) {
			respBody, err := json.Marshal(gin.H{
				"error": error.NewBadRequest("Invalid query pareameters. Verify request query parameters once."),
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Failure, While URI have empty query string", func(t *testing.T) {
		uc := uldMock.NewMockUseCaser(ctrl)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.GET("/infotax/user_login_detail/search", h.SearchUserLoginDetail)

		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "/infotax/user_login_detail/search", nil)
		assert.NoError(t, err)
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusBadRequest, result.StatusCode) {
			respBody, err := json.Marshal(gin.H{
				"error": error.NewBadRequest("Valid query pareameters not found. Verify request query parameters once."),
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Failure, DB error occurred while searching user login details", func(t *testing.T) {
		var expected []out.UserLoginDetail
		expectedError := error.NewInternal()
		uc := uldMock.NewMockUseCaser(ctrl)
		uc.EXPECT().SearchUserLoginDetail(gomock.AssignableToTypeOf(&ctx), filterMap).Return(expected, expectedError)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.GET("/infotax/user_login_detail/search", h.SearchUserLoginDetail)

		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "/infotax/user_login_detail/search"+validQuery, nil)
		assert.NoError(t, err)
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusInternalServerError, result.StatusCode) {
			respBody, err := json.Marshal(gin.H{
				"error": expectedError,
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Success, While user login details was fetched successfully", func(t *testing.T) {
		uc := uldMock.NewMockUseCaser(ctrl)
		uc.EXPECT().SearchUserLoginDetail(gomock.AssignableToTypeOf(&ctx), filterMap).Return(dtls, nil)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.GET("/infotax/user_login_detail/search", h.SearchUserLoginDetail)

		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "/infotax/user_login_detail/search"+validQuery, nil)
		assert.NoError(t, err)
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusOK, result.StatusCode) {
			respBody, err := json.Marshal(dtls)
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})
}
