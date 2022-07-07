package user_login_detail

import (
	"fmt"
	"testing"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	repoMock "www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository/mock"
)

func TestNewUseCase_SearchUserLoginDetail(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Db error occurred, while searching user login details", func(t *testing.T) {
		var ctx gin.Context
		var filterMap map[string]interface{}
		var expected []out.UserLoginDetail
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().SearchUserLoginDetail(gomock.AssignableToTypeOf(&ctx), filterMap).Return([]entity.UserLoginDetails{}, fmt.Errorf("Db error occurred"))
		uc := NewUseCase(r)
		res, err := uc.SearchUserLoginDetail(&ctx, filterMap)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, expected, res)
	})

	t.Run("Fetched user login details successfully", func(t *testing.T) {
		var ctx gin.Context
		filterMap := map[string]interface{}{
			"employee_id": "252",
			"domain_name": "s",
		}
		dtls := []entity.UserLoginDetails{
			entity.UserLoginDetails{
				EmployeeID:   "2521",
				DomainName:   "saravase",
				EmailID:      "saravase.s@ivtlinfoview.co.jp",
				Password:     "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
				UUID:         "Testdata",
				IsSignedUp:   0,
				Role:         "Payroll_User",
				EnableAccess: "No",
			},
			entity.UserLoginDetails{
				EmployeeID:   "2522",
				DomainName:   "thasarathan",
				EmailID:      "thasarathan.s@ivtlinfoview.co.jp",
				Password:     "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
				UUID:         "Testdata",
				IsSignedUp:   0,
				Role:         "Payroll_User",
				EnableAccess: "No",
			},
		}
		expected := []out.UserLoginDetail{
			out.UserLoginDetail{
				EmployeeID: "2521",
				DomainName: "saravase",
				EmailID:    "saravase.s@ivtlinfoview.co.jp",
				Password:   "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
				UUID:       "Testdata",
				IsSignedUp: 0,
				Role:       "Payroll_User",
			},
			out.UserLoginDetail{
				EmployeeID: "2522",
				DomainName: "thasarathan",
				EmailID:    "thasarathan.s@ivtlinfoview.co.jp",
				Password:   "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
				UUID:       "Testdata",
				IsSignedUp: 0,
				Role:       "Payroll_User",
			},
		}
		repoMock.NewMockUserLoginDetailRepository(ctrl)
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().SearchUserLoginDetail(gomock.AssignableToTypeOf(&ctx), filterMap).Return(dtls, nil)
		uc := NewUseCase(r)
		res, err := uc.SearchUserLoginDetail(&ctx, filterMap)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

}
