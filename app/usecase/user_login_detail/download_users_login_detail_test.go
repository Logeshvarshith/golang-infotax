package user_login_detail

import (
	"fmt"
	"path"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	repoMock "www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository/mock"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
)

func TestNewUseCase_DownloadUsersLoginDetail(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dtls := []entity.UserLoginDetails{
		entity.UserLoginDetails{
			EmployeeID:   "1",
			DomainName:   "a",
			EmailID:      "a.s@ivtlinfoview.co.jp",
			Password:     "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
			UUID:         "Testdata",
			IsSignedUp:   0,
			Role:         "Payroll_User",
			EnableAccess: "No",
		},
		entity.UserLoginDetails{
			EmployeeID:   "b",
			DomainName:   "saravase",
			EmailID:      "b.s@ivtlinfoview.co.jp",
			Password:     "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
			UUID:         "Testdata",
			IsSignedUp:   0,
			Role:         "Payroll_User",
			EnableAccess: "No",
		},
	}

	t.Run("db error occurred, while fetching users login details", func(t *testing.T) {
		var ctx gin.Context
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().GetAllUserLoginDetail(gomock.AssignableToTypeOf(&ctx)).Return([]entity.UserLoginDetails{}, fmt.Errorf("Db error occurred"))
		uc := NewUseCase(r)
		p := "tmp/test.csv"
		name, fpath, err := uc.DownloadUsersLoginDetail(&ctx, p)
		assert.Equal(t, err, error.NewInternal())
		assert.Empty(t, name)
		assert.Empty(t, fpath)
	})

	t.Run("File not found error occurred, while downloading users login details", func(t *testing.T) {
		var ctx gin.Context
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().GetAllUserLoginDetail(gomock.AssignableToTypeOf(&ctx)).Return(dtls, nil)
		uc := NewUseCase(r)
		p := "./"
		name, fpath, err := uc.DownloadUsersLoginDetail(&ctx, p)
		assert.Equal(t, err, error.NewNotFound("File", path.Base(p)))
		assert.NotEmpty(t, name)
		assert.NotEmpty(t, fpath)
	})

	t.Run("Downloaded users login details successfully", func(t *testing.T) {
		var ctx gin.Context
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().GetAllUserLoginDetail(gomock.AssignableToTypeOf(&ctx)).Return(dtls, nil)
		uc := NewUseCase(r)
		p := "tmp/test.csv"
		name, fpath, err := uc.DownloadUsersLoginDetail(&ctx, p)
		assert.Nil(t, err)
		assert.NotEmpty(t, name)
		assert.NotEmpty(t, fpath)
	})

}
