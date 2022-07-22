package employee_payroll_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"

	"github.com/jinzhu/copier"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
)

func (uc *useCase) UpdateEmployeePayrollDetail(ctx context.Context, empID string, detail in.UpdatedEmployeePayrollDetail) (updRes out.UpdatedResponse, err *error.Error) {

	id := entity.EmployeeID(empID)
	exist, rerr := uc.employeePayrollDetailRepo.CheckIfEmployeePayrollDetailExists(ctx, id)
	if !exist {
		err = error.NewNotFound("employee_id", empID)
		return
	}

	dtl, rerr := uc.employeePayrollDetailRepo.GetEmployeePayrollDetail(ctx, id)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	copier.Copy(&dtl, &detail)

	rerr = uc.employeePayrollDetailRepo.UpdateEmployeePayrollDetail(ctx, id, dtl)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	updRes = out.UpdatedResponse{
		IsUpdated: "true",
	}
	return
}
