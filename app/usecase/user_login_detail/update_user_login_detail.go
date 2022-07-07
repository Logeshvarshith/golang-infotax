package user_login_detail

import (
	"context"

	"github.com/jinzhu/copier"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// UpdateUserLoginDetail method is used to share user login detail to repository and verify the return data from repoistory.
func (uc *useCase) UpdateUserLoginDetail(ctx context.Context, empID string, detail in.UpdateUserDetail) (updRes out.UpdateResponse, err *error.Error) {

	id := entity.EmployeeID(empID)
	exist, rerr := uc.userLoginDetailRepo.CheckIfUserLoginDetailExists(ctx, id)
	if rerr != nil {
		err = error.NewInternal()
		return
	}
	if !exist {
		err = error.NewNotFound("employee_id", empID)
		return
	}

	dtl, rerr := uc.userLoginDetailRepo.GetUserLoginDetail(ctx, id)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	copier.Copy(&dtl, &detail)

	rerr = uc.userLoginDetailRepo.UpdateUserLoginDetail(ctx, id, dtl)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	updRes = out.UpdateResponse{
		IsUpdated: "true",
	}
	return
}
