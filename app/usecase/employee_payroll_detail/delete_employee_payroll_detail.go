package employee_payroll_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

// DeleteUserLoginDetail method is used to send filter detail to repository and verify the return data from repoistory.
func (uc *useCase) DeleteEmployeePayrollDetail(ctx context.Context, empID string) (delRes out.DeleteResponse, err *error.Error) {

	exist, rerr := uc.employeePayrollDetailRepo.CheckIfEmployeePayrollDetailExists(ctx, entity.EmployeeID(empID))
	if rerr != nil {
		err = error.NewInternal()
		return
	}
	if !exist {
		err = error.NewNotFound("employee_id", empID)
		return
	}

	rerr = uc.employeePayrollDetailRepo.DeleteEmployeePayrollDetail(ctx, entity.EmployeeID(empID))
	if rerr != nil {

		err = error.NewInternal()
		return
	}

	delRes = out.DeleteResponse{
		IsDeleted: "true",
	}
	return
}
