package employee_payroll_detail

import (
	"context"

	"github.com/jinzhu/copier"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

// SearchUserLoginDetail method is used to send filter detail to repository and get login user details
func (uc *useCase) SearchEmployeePayrollDetail(ctx context.Context, filterMap map[string]interface{}) (details []out.EmployeePayrollDetail, err *error.Error) {
	dtls, rerr := uc.employeePayrollDetailRepo.SearchEmployeePayrollDetail(ctx, filterMap)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	copier.Copy(&details, &dtls)

	return
}
