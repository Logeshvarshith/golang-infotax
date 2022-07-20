package employee_payroll_detail

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

func (uc *useCase) CreateEmployeePayrollDetail(ctx context.Context, payrolldetail in.CreateEmployeePayrollDetail) (SaveResponse out.PayrollDetailSaveResponse, err *error.Error) {
	var payrolldetails entity.EmployeePayrollMst

	copier.Copy(&payrolldetails, &payrolldetail)

	response := uc.employeePayrollDetailRepo.CreateEmployeePayrollDetail(ctx, payrolldetails)

	serr, ok := response.(*mysql.MySQLError)
	if response != nil && !ok {
		err = error.NewInternal()
		return
	}

	if ok && serr.Number == 1062 {
		err = error.NewConflict("employee_id", payrolldetail.EmployeeID)
		return
	}

	SaveResponse = out.PayrollDetailSaveResponse{
		IsSaved: "true",
	}
	return
}
