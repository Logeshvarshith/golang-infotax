package employee_official_detail

import (
	"context"

	"github.com/jinzhu/copier"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"

	"github.com/go-sql-driver/mysql"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/out"
)

// CreateUserLoginDetail method is used to send user login detail to repository and verify the return data from repoistory.
func (uc *useCase) CreateEmployeeOfficialDetail(ctx context.Context, detail in.CreateEmployeeOfficial) (savRes out.SavedResponse, err *error.Error) {

	var detail1 entity.EmployeeOfficialMst
	copier.Copy(&detail1, &detail)
	rerr := uc.employeeOfficialDetailRepo.CreateEmployeeOfficialDetail(ctx, detail1)

	serr, ok := rerr.(*mysql.MySQLError)
	if rerr != nil && !ok {
		err = error.NewInternal()
		return
	}

	if ok && serr.Number == 1062 {
		err = error.NewConflict("employee_id", string(detail.EmployeeID))
		return
	}

	savRes = out.SavedResponse{
		IsSaved: "true",
	}
	return
}
