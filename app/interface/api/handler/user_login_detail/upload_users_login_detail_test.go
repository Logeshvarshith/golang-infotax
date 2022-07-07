package user_login_detail

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	uldMock "www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/mock"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

func TestUserLoginDetailHandler_UploadUsersLoginDetail(t *testing.T) {
	log.MakeLogger("tmp/infotax.log", true)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var ctx gin.Context

	t.Run("Failure, While passing request payload with invalid file input", func(t *testing.T) {
		uc := uldMock.NewMockUseCaser(ctrl)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.POST("/infotax/user_login_detail/upload", h.UploadUsersLoginDetail)

		w := httptest.NewRecorder()
		file, err := os.Open("/tmp")
		assert.NoError(t, err)
		r, err := http.NewRequest("POST", "/infotax/user_login_detail/upload", file)
		assert.NoError(t, err)
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusBadRequest, result.StatusCode) {
			respBody, err := json.Marshal(gin.H{
				"error": error.NewBadRequest("Invalid request file. Verify request file."),
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Failure, While passing request payload with invalid file extension", func(t *testing.T) {
		uc := uldMock.NewMockUseCaser(ctrl)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.POST("/infotax/user_login_detail/upload", h.UploadUsersLoginDetail)

		w := httptest.NewRecorder()
		file, err := os.Open("./tmp/test.txt")
		defer file.Close()
		assert.NoError(t, err)
		buf, writer := createMultipartFormData(t, "file", file)
		writer.Close()
		r, err := http.NewRequest("POST", "/infotax/user_login_detail/upload", buf)
		assert.NoError(t, err)
		r.Header.Set("Content-Type", writer.FormDataContentType())
		router.ServeHTTP(w, r)
		result := w.Result()
		defer result.Body.Close()
		if assert.Equal(t, http.StatusBadRequest, result.StatusCode) {
			respBody, err := json.Marshal(gin.H{
				"error": error.NewBadRequest("Invalid request file extension. Verify request file extension."),
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}
	})

	t.Run("Failure, Db error occurred while doing bulk insert", func(t *testing.T) {
		var expected out.SaveResponse
		expectedErr := error.NewInternal()
		uc := uldMock.NewMockUseCaser(ctrl)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.POST("/infotax/user_login_detail/upload", h.UploadUsersLoginDetail)

		w := httptest.NewRecorder()
		file, err := os.Open("./tmp/test.csv")
		defer file.Close()
		assert.NoError(t, err)
		buf, writer := createMultipartFormData(t, "file", file)
		writer.Close()
		r, err := http.NewRequest("POST", "/infotax/user_login_detail/upload", buf)
		r.Header.Set("Content-Type", writer.FormDataContentType())
		assert.NoError(t, err)
		mf, _, err := r.FormFile("file")
		assert.NoError(t, err)
		uc.EXPECT().UploadUsersLoginDetail(gomock.AssignableToTypeOf(&ctx), mf).Return(expected, expectedErr)
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

	t.Run("Success, While file upload and bulk insertion was completed successfully", func(t *testing.T) {
		expected := out.SaveResponse{
			IsSaved: "true",
		}
		uc := uldMock.NewMockUseCaser(ctrl)
		h := NewUserLoginDetailHandler(uc)
		router := gin.Default()
		router.POST("/infotax/user_login_detail/upload", h.UploadUsersLoginDetail)

		w := httptest.NewRecorder()
		file, err := os.Open("./tmp/test.csv")
		defer file.Close()
		assert.NoError(t, err)
		buf, writer := createMultipartFormData(t, "file", file)
		writer.Close()
		r, err := http.NewRequest("POST", "/infotax/user_login_detail/upload", buf)
		r.Header.Set("Content-Type", writer.FormDataContentType())
		assert.NoError(t, err)
		mf, _, err := r.FormFile("file")
		assert.NoError(t, err)
		uc.EXPECT().UploadUsersLoginDetail(gomock.AssignableToTypeOf(&ctx), mf).Return(expected, nil)
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

func createMultipartFormData(t *testing.T, fieldName string, file *os.File) (*bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile(fieldName, file.Name())
	assert.NoError(t, err)
	_, err = io.Copy(fw, file)
	assert.NoError(t, err)
	return &b, w
}
