package user_login_detail

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	repoMock "www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository/mock"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
)

func TestUseCase_GetUserLoginRole(t *testing.T) {

	// testcase is to run in parallel
	t.Parallel()
	// create mock object
	ctrl := gomock.NewController(t)
	// cleanup mock object after all of the test methods were finished
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
		res, err := uc.GetUserLoginRole(&ctx, id)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, "", res.Role)
	})

	t.Run("Not found error occurred, while employee id exist check", func(t *testing.T) {
		var ctx gin.Context
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), inValidID).Return(false, nil)
		uc := NewUseCase(r)
		res, err := uc.GetUserLoginRole(&ctx, id)
		assert.Equal(t, err, error.NewNotFound("employee_id", id))
		assert.Equal(t, "", res.Role)
	})

	t.Run("Db error occurred, while fetching user login role", func(t *testing.T) {
		var ctx gin.Context
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), validID).Return(true, nil)
		r.EXPECT().GetUserLoginRole(gomock.AssignableToTypeOf(&ctx), validID).Return(entity.EmployeeRole(""), fmt.Errorf("Db error occurred"))
		uc := NewUseCase(r)
		res, err := uc.GetUserLoginRole(&ctx, vid)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, "", res.Role)
	})

	t.Run("Fetched user login role successfully", func(t *testing.T) {
		var ctx gin.Context
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), validID).Return(true, nil)
		r.EXPECT().GetUserLoginRole(gomock.AssignableToTypeOf(&ctx), validID).Return(entity.EmployeeRole("Payroll_User"), nil)
		uc := NewUseCase(r)
		res, err := uc.GetUserLoginRole(&ctx, vid)
		assert.Nil(t, err)
		assert.Equal(t, "Payroll_User", res.Role)
	})

}
