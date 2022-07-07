package user_login_detail

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	repoMock "www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository/mock"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

func TestNewUseCase_UpdateUserLoginDetail(t *testing.T) {

	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := ""
	inValidID := entity.EmployeeID(id)
	vid := "2521"
	validID := entity.EmployeeID(vid)

	dtl := in.UpdateUserDetail{
		DomainName: "sss",
		EmailID:    "sss@gmail.com",
	}

	t.Run("Db error occurred, while employee id exist check", func(t *testing.T) {
		var ctx gin.Context
		var expected out.UpdateResponse
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), inValidID).Return(false, fmt.Errorf("Db error occurred"))
		uc := NewUseCase(r)
		res, err := uc.UpdateUserLoginDetail(&ctx, id, dtl)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, expected, res)
	})

	t.Run("Not found error occurred, while employee id exist check", func(t *testing.T) {
		var ctx gin.Context
		var expected out.UpdateResponse
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), inValidID).Return(false, nil)
		uc := NewUseCase(r)
		res, err := uc.UpdateUserLoginDetail(&ctx, id, dtl)
		assert.Equal(t, err, error.NewNotFound("employee_id", id))
		assert.Equal(t, expected, res)
	})

	t.Run("Db error occurred, while fetching user login detail", func(t *testing.T) {
		var ctx gin.Context
		var expected out.UpdateResponse
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), validID).Return(true, nil)
		r.EXPECT().GetUserLoginDetail(gomock.AssignableToTypeOf(&ctx), validID).Return(entity.UserLoginDetails{}, fmt.Errorf("Db error occurred"))
		uc := NewUseCase(r)
		res, err := uc.UpdateUserLoginDetail(&ctx, vid, dtl)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, expected, res)
	})

	t.Run("Db error occurred, while updating user login detail", func(t *testing.T) {
		var ctx gin.Context
		var detail entity.UserLoginDetails
		var expected out.UpdateResponse
		copier.Copy(&dtl, &detail)
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), validID).Return(true, nil)
		r.EXPECT().GetUserLoginDetail(gomock.AssignableToTypeOf(&ctx), validID).Return(detail, nil)
		r.EXPECT().UpdateUserLoginDetail(gomock.AssignableToTypeOf(&ctx), validID, detail).Return(fmt.Errorf("Db error occurred"))
		uc := NewUseCase(r)
		res, err := uc.UpdateUserLoginDetail(&ctx, vid, dtl)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, expected, res)
	})

	t.Run("Fetched user login detail successfully", func(t *testing.T) {
		var ctx gin.Context
		var detail entity.UserLoginDetails
		expected := out.UpdateResponse{
			IsUpdated: "true",
		}
		copier.Copy(&dtl, &detail)
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CheckIfUserLoginDetailExists(gomock.AssignableToTypeOf(&ctx), validID).Return(true, nil)
		r.EXPECT().GetUserLoginDetail(gomock.AssignableToTypeOf(&ctx), validID).Return(detail, nil)
		r.EXPECT().UpdateUserLoginDetail(gomock.AssignableToTypeOf(&ctx), validID, detail).Return(nil)
		uc := NewUseCase(r)
		res, err := uc.UpdateUserLoginDetail(&ctx, vid, dtl)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

}
