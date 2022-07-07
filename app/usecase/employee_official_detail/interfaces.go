package employee_official_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/out"
)

type UseCaser interface {
	CreateEmployeeOfficialDetail(ctx context.Context, detail in.CreateEmployeeOfficial) (out.SavedResponse, *error.Error)
}

type useCase struct {
	employeeOfficialDetailRepo repository.EmployeeOfficialDetailRepository
}

// NewUseCase function is used to make new userCase struct.
func NewUseCase(employeeOfficialDetailRepo repository.EmployeeOfficialDetailRepository) UseCaser {
	return &useCase{
		employeeOfficialDetailRepo: employeeOfficialDetailRepo,
	}
}
