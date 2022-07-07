package mysql

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"

	"gorm.io/gorm"
)

type EmployeeOfficialDetailRepository struct {
	db *gorm.DB
}

func NewEmployeeOfficialDetailRepository(db *gorm.DB) *EmployeeOfficialDetailRepository {
	return &EmployeeOfficialDetailRepository{
		db: db,
	}
}

func (r *EmployeeOfficialDetailRepository) CreateEmployeeOfficialDetail(ctx context.Context, detail entity.EmployeeOfficialMst) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("employee_official_mst").Create(&detail)
	err = db.Error
	return
}
