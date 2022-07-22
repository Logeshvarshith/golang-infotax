package employee_payroll_detail

import (
	"context"

	"github.com/jinzhu/copier"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

// DeleteUserLoginDetail method is used to send filter detail to repository and verify the return data from repoistory.
func (uc *useCase) DeleteMultipleEmployeePayrollDetail(ctx context.Context, empID in.DeleteMultipleEmployee) (delRes out.DeleteResponse, err *error.Error) {

	var dtls entity.DeleteMultipleEmployee

	copier.Copy(&dtls, empID)

	exist, rerr := uc.employeePayrollDetailRepo.DeleteMultipleEmployeePayrollDetail(ctx, dtls)
	if !exist {
		err = error.NewNotFoundMultiple()
		return
	}
	if rerr != nil {

		err = error.NewInternal()
		return
	}

	delRes = out.DeleteResponse{
		IsDeleted: "true",
	}
	return
}
