package user_login_detail

import (
	"fmt"
	"os"
	"testing"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	repoMock "www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository/mock"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

func TestNewUseCase_UploadUsersLoginDetail(t *testing.T) {
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

	conflictDtls := []entity.UserLoginDetails{
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
			EmployeeID:   "1",
			DomainName:   "a",
			EmailID:      "a.s@ivtlinfoview.co.jp",
			Password:     "$2a$10$g.U6giqTQ3cBj5Zf/uyVpugr6WCyzqspzXMUfwlTPwN5CWC2pxPKO",
			UUID:         "Testdata",
			IsSignedUp:   0,
			Role:         "Payroll_User",
			EnableAccess: "No",
		},
	}

	t.Run("Internal error occurred, while reading content from input csv file", func(t *testing.T) {
		var ctx gin.Context
		var expected out.SaveResponse
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		file, err := os.Open("tmp/test_invalid1.csv")
		assert.NoError(t, err)

		uc := NewUseCase(r)
		res, err := uc.UploadUsersLoginDetail(&ctx, file)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, expected, res)
	})

	t.Run("Bad request error occurred, while input csv file have wrong column datatype ", func(t *testing.T) {
		var ctx gin.Context
		var expected out.SaveResponse
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		file, err := os.Open("tmp/test_invalid2.csv")
		assert.NoError(t, err)

		uc := NewUseCase(r)
		res, err := uc.UploadUsersLoginDetail(&ctx, file)
		assert.Equal(t, err, error.NewBadRequest(fmt.Sprintf("Line no: %d, CSV file contain invalid input values", 2)))
		assert.Equal(t, expected, res)
	})

	t.Run("Bad request error occurred, while doing input validation", func(t *testing.T) {
		var ctx gin.Context
		var expected out.SaveResponse
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		file, err := os.Open("tmp/test_invalid3.csv")
		assert.NoError(t, err)

		uc := NewUseCase(r)
		res, err := uc.UploadUsersLoginDetail(&ctx, file)
		assert.Equal(t, err, error.NewBadRequest(fmt.Sprintf("Line no: %d, Input validation was failed.", 3)))
		assert.Equal(t, expected, res)
	})

	t.Run("Internal error occurred, while doing bulk insertion", func(t *testing.T) {
		var ctx gin.Context
		var expected out.SaveResponse
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CreateBulkUserLoginDetail(gomock.AssignableToTypeOf(&ctx), dtls).Return(fmt.Errorf("Db error occurred"))
		file, err := os.Open("tmp/test.csv")
		assert.NoError(t, err)

		uc := NewUseCase(r)
		res, err := uc.UploadUsersLoginDetail(&ctx, file)
		assert.Equal(t, err, error.NewInternal())
		assert.Equal(t, expected, res)
	})

	t.Run("Conflict error occurred, while doing bulk insertion", func(t *testing.T) {
		var ctx gin.Context
		var expected out.SaveResponse
		merr := &mysql.MySQLError{
			Number:  1062,
			Message: "Conflict error",
		}
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CreateBulkUserLoginDetail(gomock.AssignableToTypeOf(&ctx), conflictDtls).Return(merr)
		file, err := os.Open("tmp/test_invalid4.csv")
		assert.NoError(t, err)

		uc := NewUseCase(r)
		res, err := uc.UploadUsersLoginDetail(&ctx, file)
		assert.Equal(t, err, error.NewBulkInsertConflict(merr.Message))
		assert.Equal(t, expected, res)
	})

	t.Run("Bulk insertion operation was completed successfully", func(t *testing.T) {
		var ctx gin.Context
		expected := out.SaveResponse{
			IsSaved: "true",
		}
		r := repoMock.NewMockUserLoginDetailRepository(ctrl)
		r.EXPECT().CreateBulkUserLoginDetail(gomock.AssignableToTypeOf(&ctx), dtls).Return(nil)
		file, err := os.Open("tmp/test.csv")
		assert.NoError(t, err)

		uc := NewUseCase(r)
		res, err := uc.UploadUsersLoginDetail(&ctx, file)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

}
