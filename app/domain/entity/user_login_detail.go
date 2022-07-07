package entity

type EmployeeID string

type EmployeeRole string

type FilterMap map[string]interface{}

type UserLoginDetails struct {
	EmployeeID   string
	DomainName   string
	EmailID      string
	Password     string
	UUID         string
	IsSignedUp   int `gorm:"column:isSignedup"`
	EnableAccess string
	Role         string
}
