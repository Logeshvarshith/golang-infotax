package mysql

import (
	"context"
	"fmt"

	"github.com/thoas/go-funk"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"

	"gorm.io/gorm"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

type UserLoginDetailRepository struct {
	db *gorm.DB
}

// NewUserLoginDetailRepository function is used to make UserLoginDetailRepository struct.
func NewUserLoginDetailRepository(db *gorm.DB) *UserLoginDetailRepository {
	return &UserLoginDetailRepository{
		db: db,
	}
}

// CheckIfUserLoginDetailExists method is used to verify whether given employee_id login detail exists or not.
func (r *UserLoginDetailRepository) CheckIfUserLoginDetailExists(ctx context.Context, empID entity.EmployeeID) (exist bool, err error) {
	var dtl out.UserLoginDetail
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").First(&dtl, empID)
	err = db.Error
	if dtl.EmailID != "" {
		exist = true
	}
	return
}

// GetUserLoginDetail method is used to fetch record from user_login_details table based on the filter condition.
func (r *UserLoginDetailRepository) GetUserLoginDetail(ctx context.Context, id entity.EmployeeID) (dtl entity.UserLoginDetails, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Select("employee_id", "domain_name", "email_id", "password", "uuid", "isSignedup", "role").Where("employee_id=?", id).Scan(&dtl)
	err = db.Error
	return
}

// GetUserLoginRole method is used to fetch role from user_login_details table based on the filter condition.
func (r *UserLoginDetailRepository) GetUserLoginRole(ctx context.Context, id entity.EmployeeID) (role entity.EmployeeRole, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Select("role").Where("employee_id=?", id).Scan(&role)
	err = db.Error
	return
}

// GetAllUserLoginDetail method is used to fetch all the user login details from user_login_details table.
func (r *UserLoginDetailRepository) GetAllUserLoginDetail(ctx context.Context) (dtls []entity.UserLoginDetails, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Find(&dtls)
	err = db.Error
	return

}

// DeleteUserLoginDetail method is used to delete user login detail from user_login_details table based on the filter condition.
func (r *UserLoginDetailRepository) DeleteUserLoginDetail(ctx context.Context, id entity.EmployeeID) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Where("employee_id=?", id).Delete(&entity.UserLoginDetails{})
	err = db.Error
	return
}

// CreateUserLoginDetail method is used to create user login detail in user_login_details table.
func (r *UserLoginDetailRepository) CreateUserLoginDetail(ctx context.Context, dtl entity.UserLoginDetails) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Create(&dtl)
	err = db.Error
	return
}

// UpdateUserLoginDetail method is used to update user login detail from user_login_details table.
func (r *UserLoginDetailRepository) UpdateUserLoginDetail(ctx context.Context, id entity.EmployeeID, dtl entity.UserLoginDetails) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Where("employee_id=?", id).Updates(dtl)
	err = db.Error
	return
}

// SearchUserLoginDetail method is used to get user login details from user_login_details table based on filterMap conditions.
func (r *UserLoginDetailRepository) SearchUserLoginDetail(ctx context.Context, filterMap entity.FilterMap) (dtls []entity.UserLoginDetails, err error) {
	tx := r.db.WithContext(ctx)

	db := tx.Table("user_login_details")
	funk.ForEach(filterMap, func(key string, value interface{}) {
		if key != "isSignedup" {
			col := fmt.Sprintf("%s LIKE ?", key)
			colVal := fmt.Sprintf("%%%v%%", value)
			db = db.Where(col, colVal)
		} else {
			db = db.Where(key+"= ?", value)
		}
	})
	db.Find(&dtls)
	err = db.Error
	return
}

// CreateBulkUserLoginDetail method is used to create bulk user login detail in user_login_details table.
func (r *UserLoginDetailRepository) CreateBulkUserLoginDetail(ctx context.Context, dtls []entity.UserLoginDetails) (err error) {
	tx := r.db.WithContext(ctx).Begin()
	err = tx.Table("user_login_details").CreateInBatches(dtls, 500).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
