package in

type CreateEmployeeOfficial struct {
	EmployeeID     int64  `json:"employee_id" validate:"required"`
	ProjectID      int64  `json:"project_id"`
	DomainName     string `json:"domain_name"`
	GradeID        int64  `json:"grade_id"`
	OfficialMailID string `json:"official_mail_id" validate:"email"`
	DateOfJoining  string `json:"date_of_joining"`
	Location       string `json:"location"`
	FloorNumber    string `json:"floor_number"`
	SeatNumber     string `json:"seat_number"`
}

type UpdatedEmployeeOfficial struct {
	ProjectID      int64  `json:"project_id"`
	DomainName     string `json:"domain_name"`
	GradeID        int64  `json:"grade_id"`
	OfficialMailID string `json:"official_mail_id" validate:"email"`
	DateOfJoining  string `json:"date_of_joining"`
	Location       string `json:"location"`
	FloorNumber    string `json:"floor_number"`
	SeatNumber     string `json:"seat_number"`
}
