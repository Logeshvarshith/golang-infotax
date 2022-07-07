package user_login_detail

import (
	"fmt"
	"testing"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	repoMock "www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository/mock"
)

func TestNewUseCase_CreateUserLoginDetail(t *testing.T) {
	var ctx gin.Context
	var dtl entity.UserLoginDetails
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	detail := in.CreateUserDetail{
		EmployeeID: "6548",
		DomainName: "primz",
		EmailID:    "primz.o@ivtlinfoview.co.jp",
		Role:       "Payroll_User",
	}
	copier.Copy(&dtl, &detail)

	t.Run("Db error occurred, while creating user login detail", func(t *testing.T) {
		var expected out.SaveResponse
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CreateUserLoginDetail(gomock.AssignableToTypeOf(&ctx), dtl).Return(fmt.Errorf("Db error occurred"))
		uc := NewUseCase(r)
		res, err := uc.CreateUserLoginDetail(&ctx, detail)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, expected, res)
	})

	t.Run("Conflict error occrred, while creating user login detail", func(t *testing.T) {
		var expected out.SaveResponse
		copier.Copy(&dtl, &detail)
		merr := &mysql.MySQLError{
			Number:  1062,
			Message: "Conflict error",
		}
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CreateUserLoginDetail(gomock.AssignableToTypeOf(&ctx), dtl).Return(merr)
		uc := NewUseCase(r)
		res, err := uc.CreateUserLoginDetail(&ctx, detail)
		assert.Equal(t, err, error.NewConflict("employee_id", detail.EmployeeID))
		assert.Equal(t, expected, res)
	})

	t.Run("Created user login details successfully", func(t *testing.T) {
		expected := out.SaveResponse{
			IsSaved: "true",
		}
		copier.Copy(&dtl, &detail)
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CreateUserLoginDetail(gomock.AssignableToTypeOf(&ctx), dtl).Return(nil)
		uc := NewUseCase(r)
		res, err := uc.CreateUserLoginDetail(&ctx, detail)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})
}
