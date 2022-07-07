package in

type CreateEmployeeOfficial struct {
	EmployeeID     int64  `json:"employee_id" validate:"required"`
	ProjectID      int64  `json:"project_id" validate:"required"`
	DomainName     string `json:"domain_name" validate:"required"`
	GradeID        int64  `json:"grade_id"`
	OfficialMailID string `json:"official_mail_id" validate:"required,email"`
	DateOfJoining  string `json:"date_of_joining"`
	Location       string `json:"location"`
	FloorNumber    string `json:"floor_number"`
	SeatNumber     string `json:"seat_number"`
}
