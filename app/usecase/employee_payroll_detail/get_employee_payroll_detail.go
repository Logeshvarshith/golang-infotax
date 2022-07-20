package employee_payroll_detail

import (
	"context"

	"github.com/jinzhu/copier"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

func (uc *useCase) GetAllEmployeePayrollDetail(ctx context.Context) (employeepayrolldetail []out.EmployeePayrollDetail, err *error.Error) {
	employeepayrolldetail1, rerr := uc.employeePayrollDetailRepo.GetAllEmployeePayrollDetail(ctx)

	if rerr != nil {
		err = error.NewInternal()
		return
	}

	copier.Copy(&employeepayrolldetail, &employeepayrolldetail1)

	return
}
