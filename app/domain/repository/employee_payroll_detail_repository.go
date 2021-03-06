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
	DeleteMultipleEmployeePayrollDetail(ctx context.Context, empID entity.DeleteMultipleEmployee) (bool, error)
	UpdateEmployeePayrollDetail(ctx context.Context, empID entity.EmployeeID, detail entity.EmployeePayrollMst) error
	GetEmployeePayrollDetail(ctx context.Context, empID entity.EmployeeID) (entity.EmployeePayrollMst, error)
	SearchEmployeePayrollDetail(ctx context.Context, filterMap entity.FilterMap) ([]entity.EmployeePayrollMst, error)
}
