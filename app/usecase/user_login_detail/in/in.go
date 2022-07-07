package in

type CreateUserDetail struct {
	EmployeeID   string `json:"employee_id" validate:"required"`
	DomainName   string `json:"domain_name" validate:"required"`
	EmailID      string `json:"email_id" validate:"required,email"`
	Password     string `json:"-" validate:"-"`
	UUID         string `json:"-" validate:"-"`
	IsSignedUp   int    `json:"-" validate:"-" gorm:"column:isSignedup"`
	EnableAccess string `json:"-" validate:"-"`
	Role         string `json:"role" validate:"required,oneof=Payroll_User Accounts_Team Payroll_Team"`
} // @name CreateUserDetail

type UpdateUserDetail struct {
	DomainName string `json:"domain_name" validate:"required"`
	EmailID    string `json:"email_id" validate:"omitempty,email"`
	Role       string `json:"role" validate:"omitempty,oneof=Payroll_User Accounts_Team Payroll_Team"`
} // @name UpdateUserDetail

type EmployeeIDFilter struct {
	EmployeeID string `json:"employee_id" validate:"required"`
} // @name EmployeeIDFilter

type UploadUserDetail struct {
	EmployeeID   string `json:"employee_id" validate:"required"`
	DomainName   string `json:"domain_name" validate:"required"`
	EmailID      string `json:"email_id" validate:"required,email"`
	Password     string `json:"-" validate:"required"`
	UUID         string `json:"-" validate:"required"`
	IsSignedUp   int    `json:"-" validate:"-" gorm:"column:isSignedup"`
	EnableAccess string `json:"-" validate:"required"`
	Role         string `json:"role" validate:"required,oneof=Payroll_User Accounts_Team Payroll_Team"`
} // @name UploadUserDetail
