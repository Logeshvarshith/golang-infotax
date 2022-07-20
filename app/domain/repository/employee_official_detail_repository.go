package repository

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
)

type EmployeeOfficialDetailRepository interface {
	CreateEmployeeOfficialDetail(ctx context.Context, detail entity.EmployeeOfficialMst) error
	CheckIfEmployeeOfficialDetailExists(ctx context.Context, empID entity.EmployeeID) (bool, error)
	GetEmployeeOfficialDetail(ctx context.Context, empID entity.EmployeeID) (entity.EmployeeOfficialMst, error)
	UpdateEmployeeOfficialDetail(ctx context.Context, id entity.EmployeeID, dtl entity.EmployeeOfficialMst) error
}
