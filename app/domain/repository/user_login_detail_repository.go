package repository

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
)

// UserLoginDetailRepository
type UserLoginDetailRepository interface {
	GetUserLoginDetail(ctx context.Context, id entity.EmployeeID) (entity.UserLoginDetails, error)
	GetUserLoginRole(ctx context.Context, id entity.EmployeeID) (entity.EmployeeRole, error)
	GetAllUserLoginDetail(ctx context.Context) ([]entity.UserLoginDetails, error)
	DeleteUserLoginDetail(ctx context.Context, id entity.EmployeeID) error
	CreateUserLoginDetail(ctx context.Context, dtl entity.UserLoginDetails) error
	UpdateUserLoginDetail(ctx context.Context, id entity.EmployeeID, dtl entity.UserLoginDetails) error
	CheckIfUserLoginDetailExists(ctx context.Context, empID entity.EmployeeID) (bool, error)
	SearchUserLoginDetail(ctx context.Context, filterMap entity.FilterMap) ([]entity.UserLoginDetails, error)
	CreateBulkUserLoginDetail(ctx context.Context, dtls []entity.UserLoginDetails) error
}
