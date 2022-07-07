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
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

func TestUserLoginDetailHandler_DownloadUserLoginDetailTemplate(t *testing.T) {
	log.MakeLogger("tmp/infotax.log", true)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var ctx gin.Context

	actualName := "outfile/user_login_detail_template.csv"
	t.Run("Failure, While giving invalid file path", func(t *testing.T) {
		expectedErr := error.NewNotFound("File", "")
		useCaser := uldMock.NewMockUseCaser(ctrl)
		useCaser.EXPECT().DownloadUserLoginDetailTemplate(gomock.AssignableToTypeOf(&ctx), actualName).Return("", "", expectedErr)
		h := NewUserLoginDetailHandler(useCaser)
		router := gin.Default()
		router.GET("/infotax/user_login_detail/download/template", h.DownloadUserLoginDetailTemplate)

		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "/infotax/user_login_detail/download/template", nil)
		assert.NoError(t, err)
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

	t.Run("Success, While giving valid file path", func(t *testing.T) {
		path := "./tmp/test_template.csv"
		name := "test.csv"
		useCaser := uldMock.NewMockUseCaser(ctrl)
		useCaser.EXPECT().DownloadUserLoginDetailTemplate(gomock.AssignableToTypeOf(&ctx), actualName).Return(name, path, nil)
		h := NewUserLoginDetailHandler(useCaser)
		router := gin.Default()
		router.GET("/infotax/user_login_detail/download/template", h.DownloadUserLoginDetailTemplate)

		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "/infotax/user_login_detail/download/template", nil)
		assert.NoError(t, err)
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusOK, result.StatusCode) {
			assert.NoError(t, err)
		}
	})

}
