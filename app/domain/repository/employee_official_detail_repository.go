package repository

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"
)

type EmployeeOfficialDetailRepository interface {
	CreateEmployeeOfficialDetail(ctx context.Context, detail entity.EmployeeOfficialMst) error
}
