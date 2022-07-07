package registry

import (
	"github.com/google/wire"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail"
)

var (
	UserLoginDetailUseCaseSet = wire.NewSet(
		repositorySet,
		user_login_detail.NewUseCase,
	)
)

var (
	EmployeeOfficialDetailUseCaseSet = wire.NewSet(
		employeeRepositorySet,
		employee_official_detail.NewUseCase,
	)
)
