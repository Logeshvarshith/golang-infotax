package user_login_detail

import (
	"context"

	"github.com/jinzhu/copier"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"

	"github.com/go-sql-driver/mysql"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// CreateUserLoginDetail method is used to send user login detail to repository and verify the return data from repoistory.
func (uc *useCase) CreateUserLoginDetail(ctx context.Context, detail in.CreateUserDetail) (savRes out.SaveResponse, err *error.Error) {

	var dtl entity.UserLoginDetails
	copier.Copy(&dtl, &detail)
	rerr := uc.userLoginDetailRepo.CreateUserLoginDetail(ctx, dtl)

	serr, ok := rerr.(*mysql.MySQLError)
	if rerr != nil && !ok {
		err = error.NewInternal()
		return
	}

	if ok && serr.Number == 1062 {
		err = error.NewConflict("employee_id", detail.EmployeeID)
		return
	}

	savRes = out.SaveResponse{
		IsSaved: "true",
	}
	return
}
