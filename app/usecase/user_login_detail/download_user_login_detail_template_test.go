package user_login_detail

import (
	"path"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	repoMock "www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository/mock"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
)

func TestNewUseCase_DownloadUserLoginDetailTemplate(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("File not found error occurred, while downloading users login details template", func(t *testing.T) {
		var ctx gin.Context
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		uc := NewUseCase(r)
		p := "./"
		name, fpath, err := uc.DownloadUserLoginDetailTemplate(&ctx, p)
		assert.Equal(t, err, error.NewNotFound("File", path.Base(p)))
		assert.NotEmpty(t, name)
		assert.NotEmpty(t, fpath)
	})

	t.Run("Downloaded users login details template successfully", func(t *testing.T) {
		var ctx gin.Context
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		uc := NewUseCase(r)
		p := "tmp/test_template.csv"
		name, fpath, err := uc.DownloadUserLoginDetailTemplate(&ctx, p)
		assert.Nil(t, err)
		assert.NotEmpty(t, name)
		assert.NotEmpty(t, fpath)
	})

}
