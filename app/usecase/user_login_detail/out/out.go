package out

type UserLoginDetail struct {
	EmployeeID string `json:"employee_id"`
	DomainName string `json:"domain_name"`
	EmailID    string `json:"email_id"`
	Password   string `json:"password"`
	UUID       string `json:"uuid"`
	IsSignedUp int    `json:"isSignedup"`
	Role       string `json:"role"`
} // @name UserLoginDetail

type UserLoginRole struct {
	Role string `json:"role"`
} // @name UserLoginRole

type DeleteResponse struct {
	IsDeleted string `json:"isDeleted"`
} // @name DeleteResponse

type SaveResponse struct {
	IsSaved string `json:"isSaved"`
} // @name SaveResponse

type UpdateResponse struct {
	IsUpdated string `json:"isUpdated"`
} // @name UpdateResponse
