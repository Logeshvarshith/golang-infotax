package user_login_detail

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/jinzhu/copier"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// GetUserLoginDetail method is used to send data to userLoginDetail repository and verify the return data from repoistory.
func (uc *useCase) GetUserLoginDetail(ctx context.Context, empID string) (userLoginDetail out.UserLoginDetail, err *error.Error) {

	exist, rerr := uc.userLoginDetailRepo.CheckIfUserLoginDetailExists(ctx, entity.EmployeeID(empID))
	if rerr != nil && !(errors.Is(rerr, gorm.ErrRecordNotFound)) {
		err = error.NewInternal()
		return
	}
	if !exist {
		err = error.NewNotFound("employee_id", empID)
		return
	}

	dtl, rerr := uc.userLoginDetailRepo.GetUserLoginDetail(ctx, entity.EmployeeID(empID))
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	copier.Copy(&userLoginDetail, &dtl)

	return
}
