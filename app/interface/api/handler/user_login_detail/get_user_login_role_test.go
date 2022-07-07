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

func TestUserLoginDetailHandler_GetUserLoginRole(t *testing.T) {
	log.MakeLogger("tmp/infotax.log", true)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var ctx gin.Context
	eid := " "
	vid := "2521"

	t.Run("Failure, While give invalid employee id", func(t *testing.T) {
		var expected out.UserLoginRole
		expectedErr := error.NewNotFound("employee_id", eid)
		useCaser := uldMock.NewMockUseCaser(ctrl)
		useCaser.EXPECT().GetUserLoginRole(gomock.AssignableToTypeOf(&ctx), eid).Return(expected, expectedErr)
		h := NewUserLoginDetailHandler(useCaser)
		router := gin.Default()
		router.GET("/infotax/user_login_detail/role/:emp_id", h.GetUserLoginRole)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/infotax/user_login_detail/role/"+eid, nil)
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusNotFound, result.StatusCode) {
			respBody, err := json.Marshal(gin.H{
				"error": expectedErr,
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Success, While give valid employee id", func(t *testing.T) {
		expected := out.UserLoginRole{
			Role: "Payroll_User",
		}
		useCaser := uldMock.NewMockUseCaser(ctrl)
		useCaser.EXPECT().GetUserLoginRole(gomock.AssignableToTypeOf(&ctx), vid).Return(expected, nil)
		h := NewUserLoginDetailHandler(useCaser)
		router := gin.Default()
		router.GET("/infotax/user_login_detail/role/:emp_id", h.GetUserLoginRole)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/infotax/user_login_detail/role/"+vid, nil)
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
