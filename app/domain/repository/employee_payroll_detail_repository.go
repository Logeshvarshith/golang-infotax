package repository

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
)

// UserLoginDetailRepository
type EmployeePayrollDetailRepository interface {
	CreateEmployeePayrollDetail(ctx context.Context, payrolldetail entity.EmployeePayrollMst) error
	GetAllEmployeePayrollDetail(ctx context.Context) ([]entity.EmployeePayrollMst, error)
	DeleteEmployeePayrollDetail(ctx context.Context, empID entity.EmployeeID) error
	CheckIfEmployeePayrollDetailExists(ctx context.Context, empID entity.EmployeeID) (bool, error)
}
