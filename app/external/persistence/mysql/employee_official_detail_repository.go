package mysql

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/out"

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

func (r *EmployeeOfficialDetailRepository) CheckIfEmployeeOfficialDetailExists(ctx context.Context, empID entity.EmployeeID) (exist bool, err error) {
	var dtl out.UpdateEmployeeOfficial
	tx := r.db.WithContext(ctx)
	db := tx.Table("employee_official_mst").First(&dtl, empID)
	err = db.Error
	if dtl.OfficialMailID != "" {
		exist = true
	}
	return
}

func (r *EmployeeOfficialDetailRepository) GetEmployeeOfficialDetail(ctx context.Context, id entity.EmployeeID) (dtl entity.EmployeeOfficialMst, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("employee_official_mst").Select("employee_id", "project_id", "domain_name", "grade_id", "official_mail_id", "date_of_joining", "location", "floor_number", "seat_number").Where("employee_id=?", id).Scan(&dtl)
	err = db.Error
	return
}

func (r *EmployeeOfficialDetailRepository) UpdateEmployeeOfficialDetail(ctx context.Context, id entity.EmployeeID, dtl entity.EmployeeOfficialMst) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("employee_official_mst").Where("employee_id=?", id).Updates(dtl)
	err = db.Error
	return
}
