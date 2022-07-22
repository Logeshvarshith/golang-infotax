package in

type CreateEmployeePayrollDetail struct {
	EmployeeID        string `json:"employee_id" validate:"required"`
	PanNumber         string `json:"pan_number"`
	UanNumber         int64  `json:"uan_number"`
	BankAccountNumber int64  `json:"bank_account_number"`
	BankIfscCode      string `json:"bank_ifsc_code"`
	PassportNumber    string `json:"passport_number"`
	PfAccountNumber   string `json:"pf_account_number"`
	TaxRegime         string `json:"tax_regime"`
	EffectiveFrom     string `json:"effective_from"`
	EpsAccountNumber  string `json:"eps_account_number"`
	PrAccountNumber   string `json:"pr_account_number"`
	EsiNumber         string `json:"esi_number"`
}

type UpdatedEmployeePayrollDetail struct {
	PanNumber         string `json:"pan_number"`
	UanNumber         int64  `json:"uan_number"`
	BankAccountNumber int64  `json:"bank_account_number"`
	BankIfscCode      string `json:"bank_ifsc_code"`
	PassportNumber    string `json:"passport_number"`
	PfAccountNumber   string `json:"pf_account_number"`
	TaxRegime         string `json:"tax_regime"`
	EffectiveFrom     string `json:"effective_from"`
	EpsAccountNumber  string `json:"eps_account_number"`
	PrAccountNumber   string `json:"pr_account_number"`
	EsiNumber         string `json:"esi_number"`
}

type DeleteMultipleEmployee struct {
	EmployeeID []string `json:"employee_id"`
}
