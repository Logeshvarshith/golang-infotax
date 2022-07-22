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

func (ep *EmployeePayrollDetailRepository) CheckIfEmployeePayrollDetailExists(ctx context.Context, empID entity.EmployeeID) (exist bool, err error) {
	var dtl out.EmployeePayrollDetail
	tx := ep.db.WithContext(ctx)
	db := tx.Table("employee_payroll_mst").First(&dtl, empID)
	err = db.Error
	if dtl.EmployeeID != "" {
		exist = true
	}
	return
}

func (ep *EmployeePayrollDetailRepository) GetEmployeePayrollDetail(ctx context.Context, id entity.EmployeeID) (dtl entity.EmployeePayrollMst, err error) {
	tx := ep.db.WithContext(ctx)
	db := tx.Table("employee_payroll_mst").Select("employee_id", "pan_number", "uan_number", "bank_account_number", "bank_ifsc_code", "passport_number", "pf_account_number", "tax_regime", "effective_from", "eps_account_number", "pr_account_number", "esi_number").Where("employee_id=?", id).Scan(&dtl)
	err = db.Error
	return
}

func (ep *EmployeePayrollDetailRepository) UpdateEmployeePayrollDetail(ctx context.Context, id entity.EmployeeID, dtl entity.EmployeePayrollMst) (err error) {
	tx := ep.db.WithContext(ctx)
	db := tx.Table("employee_payroll_mst").Where("employee_id=?", id).Updates(dtl)
	err = db.Error
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

func (ep *EmployeePayrollDetailRepository) DeleteEmployeePayrollDetail(ctx context.Context, id entity.EmployeeID) (err error) {
	tx := ep.db.WithContext(ctx)
	db := tx.Table("employee_payroll_mst").Where("employee_id=?", id).Delete(&entity.EmployeePayrollMst{})
	err = db.Error
	return
}
