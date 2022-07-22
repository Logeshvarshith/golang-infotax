package employee_official_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/out"

	"github.com/jinzhu/copier"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/in"
)

func (uc *useCase) UpdateEmployeeOfficialDetail(ctx context.Context, empID string, detail in.UpdatedEmployeeOfficial) (updRes out.UpdatedResponse, err *error.Error) {

	id := entity.EmployeeID(empID)
	exist, rerr := uc.employeeOfficialDetailRepo.CheckIfEmployeeOfficialDetailExists(ctx, id)
	if !exist {
		err = error.NewNotFound("employee_id", empID)
		return
	}
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	dtl, rerr := uc.employeeOfficialDetailRepo.GetEmployeeOfficialDetail(ctx, id)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	copier.Copy(&dtl, &detail)

	rerr = uc.employeeOfficialDetailRepo.UpdateEmployeeOfficialDetail(ctx, id, dtl)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	updRes = out.UpdatedResponse{
		IsUpdated: "true",
	}
	return
}
