package user_login_detail

import (
	"fmt"
	"testing"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	repoMock "www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository/mock"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
)

func TestNewUseCase_GetUserLoginDetail(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := ""
	inValidID := entity.EmployeeID(id)
	vid := "2521"
	validID := entity.EmployeeID(vid)

	t.Run("Db error occurred, while employee id exist check", func(t *testing.T) {
		var ctx gin.Context
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), inValidID).Return(false, fmt.Errorf("Db error occurred"))
		uc := NewUseCase(r)
		res, err := uc.GetUserLoginDetail(&ctx, id)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, "", res.Role)
	})

	t.Run("Not found error occurred, while employee id exist check", func(t *testing.T) {
		var ctx gin.Context
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), inValidID).Return(false, nil)
		uc := NewUseCase(r)
		res, err := uc.GetUserLoginDetail(&ctx, id)
		assert.Equal(t, err, error.NewNotFound("employee_id", id))
		assert.Equal(t, "", res.Role)
	})

	t.Run("Db error occurred, while fetching user login detail", func(t *testing.T) {
		var ctx gin.Context
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), validID).Return(true, nil)
		r.EXPECT().GetUserLoginDetail(gomock.AssignableToTypeOf(&ctx), validID).Return(entity.UserLoginDetails{}, fmt.Errorf("Db error occurred"))
		uc := NewUseCase(r)
		res, err := uc.GetUserLoginDetail(&ctx, vid)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, out.UserLoginDetail{}, res)
	})

	t.Run("Fetched user login detail successfully", func(t *testing.T) {
		var ctx gin.Context
		dtl := entity.UserLoginDetails{
			EmployeeID:   "2521",
			DomainName:   "saravase",
			EmailID:      "saravanakumar.s@ivtlinfoview.co.jp",
			Password:     "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
			UUID:         "Testdata",
			IsSignedUp:   0,
			Role:         "Payroll_User",
			EnableAccess: "No",
		}
		expected := out.UserLoginDetail{
			EmployeeID: "2521",
			DomainName: "saravase",
			EmailID:    "saravanakumar.s@ivtlinfoview.co.jp",
			Password:   "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
			UUID:       "Testdata",
			IsSignedUp: 0,
			Role:       "Payroll_User",
		}
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), validID).Return(true, nil)
		r.EXPECT().GetUserLoginDetail(gomock.AssignableToTypeOf(&ctx), validID).Return(dtl, nil)
		uc := NewUseCase(r)
		res, err := uc.GetUserLoginDetail(&ctx, vid)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

}
