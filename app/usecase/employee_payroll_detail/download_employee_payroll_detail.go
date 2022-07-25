package employee_payroll_detail

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"path"

	"github.com/jinzhu/copier"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

// DownloadUsersLoginDetail method is used to return users login details in CSV format.
func (uc *useCase) DownloadEmployeePayrollDetail(ctx context.Context, fp string) (fileName, filePath string, err *error.Error) {
	var details []out.EmployeePayrollDetail
	dtls, rerr := uc.employeePayrollDetailRepo.GetAllEmployeePayrollDetail(ctx)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	copier.Copy(&details, &dtls)
	filePath = fp
	filetmp, ferr := os.Create(filePath)
	defer filetmp.Close()
	fileName = path.Base(filePath)
	if filePath == "" || fileName == "" || ferr != nil {
		err = error.NewNotFound("File", fileName)
		return
	}

	csvWriter := csv.NewWriter(filetmp)
	defer csvWriter.Flush()

	dtlHeader := []string{"employee_id", "pan_number", "uan_number", "bank_account_number", "bank_ifsc_code",
		"passport_number", "pf_account_number", "tax_regime", "effective_from", "eps_account_number", "pr_account_number", "esi_number"}
	csvWriter.Write(dtlHeader)

	for _, detail := range details {
		dtl := []string{detail.EmployeeID, detail.PanNumber, fmt.Sprintf("%v", detail.UanNumber), fmt.Sprintf("%v", detail.BankAccountNumber), detail.BankIfscCode,
			detail.PassportNumber, detail.PfAccountNumber, detail.TaxRegime, detail.EffectiveFrom, detail.EpsAccountNumber, detail.PrAccountNumber, detail.EsiNumber}
		csvWriter.Write(dtl)
	}

	return
}
