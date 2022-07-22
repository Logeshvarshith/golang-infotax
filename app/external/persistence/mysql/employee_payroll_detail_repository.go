package mysql

import (
	"context"

	"gorm.io/gorm"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

type EmployeePayrollDetailRepository struct {
	db *gorm.DB
}

func NewEmployeePayrollDetailRepository(db *gorm.DB) *EmployeePayrollDetailRepository {
	return &EmployeePayrollDetailRepository{
		db: db,
	}
}

func (r *EmployeePayrollDetailRepository) CheckIfEmployeePayrollDetailExists(ctx context.Context, empID entity.EmployeeID) (exist bool, err error) {
	var dtl out.EmployeePayrollDetail
	tx := r.db.WithContext(ctx)
	db := tx.Table("employee_payroll_mst").First(&dtl, empID)
	err = db.Error
	if dtl.PanNumber != "" {
		exist = true
	}
	return
}

func (ep *EmployeePayrollDetailRepository) GetAllEmployeePayrollDetail(ctx context.Context) (employeepayrolldetail []entity.EmployeePayrollMst, err error) {
	tx := ep.db.WithContext(ctx)
	db := tx.Table("employee_payroll_mst").Find(&employeepayrolldetail)
	err = db.Error
	return
}

func (ep *EmployeePayrollDetailRepository) CreateEmployeePayrollDetail(ctx context.Context, payrolldetail entity.EmployeePayrollMst) (err error) {
	tx := ep.db.WithContext(ctx)
	db := tx.Table("employee_payroll_mst").Create(&payrolldetail)
	err = db.Error
	return
}

func (r *EmployeePayrollDetailRepository) DeleteEmployeePayrollDetail(ctx context.Context, id entity.EmployeeID) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("employee_payroll_mst").Where("employee_id=?", id).Delete(&entity.EmployeePayrollMst{})
	err = db.Error
	return
}

func (r *EmployeePayrollDetailRepository) DeleteMultipleEmployeePayrollDetail(ctx context.Context, id entity.DeleteMultipleEmployee) (exist bool, err error) {
	tx := r.db.WithContext(ctx)

	db := tx.Table("employee_payroll_mst").Where("employee_id IN(?)", id.EmployeeID).Delete(&entity.EmployeePayrollMst{})
	err = db.Error
	if db.RowsAffected == 0 {
		exist = false
	}
	exist = true
	return
}
