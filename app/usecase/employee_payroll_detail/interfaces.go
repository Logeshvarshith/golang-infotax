package employee_payroll_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
)

type UseCaser interface {
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
