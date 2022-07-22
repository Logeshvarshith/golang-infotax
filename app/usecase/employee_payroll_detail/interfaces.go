package employee_payroll_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

type UseCaser interface {
	CreateEmployeePayrollDetail(ctx context.Context, payrolldetail in.CreateEmployeePayrollDetail) (out.PayrollDetailSaveResponse, *error.Error)
	GetAllEmployeePayrollDetail(ctx context.Context) ([]out.EmployeePayrollDetail, *error.Error)
	DeleteEmployeePayrollDetail(ctx context.Context, empID string) (out.DeleteResponse, *error.Error)
	UpdateEmployeePayrollDetail(ctx context.Context, empID string, detail in.UpdatedEmployeePayrollDetail) (out.UpdatedResponse, *error.Error)
	DownloadEmployeePayrollDetailTemplate(ctx context.Context, filePath string) (string, string, *error.Error)
}

type useCase struct {
	employeePayrollDetailRepo repository.EmployeePayrollDetailRepository
}

// NewUseCase function is used to make new userCase struct.
func NewUseCase(employeePayrollDetailRepo repository.EmployeePayrollDetailRepository) UseCaser {
	return &useCase{
		employeePayrollDetailRepo: employeePayrollDetailRepo,
	}
}
