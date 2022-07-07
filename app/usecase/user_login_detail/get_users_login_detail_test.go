package user_login_detail

import (
	"fmt"
	"testing"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	repoMock "www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository/mock"
)

func TestNewUseCase_GetAllUserLoginDetail(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("db error occurred, while fetching users login details", func(t *testing.T) {
		var ctx gin.Context
		var expected []out.UserLoginDetail
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().GetAllUserLoginDetail(gomock.AssignableToTypeOf(&ctx)).Return([]entity.UserLoginDetails{}, fmt.Errorf("Db error occurred"))
		uc := NewUseCase(r)
		res, err := uc.GetAllUserLoginDetail(&ctx)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, expected, res)
	})

	t.Run("Fetched users login details successfully", func(t *testing.T) {
		var ctx gin.Context
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
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().GetAllUserLoginDetail(gomock.AssignableToTypeOf(&ctx)).Return(dtls, nil)
		uc := NewUseCase(r)
		res, err := uc.GetAllUserLoginDetail(&ctx)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
		assert.Equal(t, len(expected), len(res))
	})

}
