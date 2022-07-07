package employee_payroll_detail

import (
	"context"
	"encoding/csv"
	"os"
	"path"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
)

// DownloadUserLoginDetailTemplate method is used to return users login details template in CSV format.
func (uc *useCase) DownloadEmployeePayrollDetailTemplate(ctx context.Context, fp string) (fileName, filePath string, err *error.Error) {

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
	return
}
