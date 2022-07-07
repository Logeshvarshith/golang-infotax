package mysql

import (
	"gorm.io/gorm"
)

type EmployeePayrollDetailRepository struct {
	db *gorm.DB
}

func NewEmployeePayrollDetailRepository(db *gorm.DB) *EmployeePayrollDetailRepository {
	return &EmployeePayrollDetailRepository{
		db: db,
	}
}
