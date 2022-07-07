package user_login_detail

import (
	"context"
	"encoding/csv"
	"os"
	"path"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
)

// DownloadUserLoginDetailTemplate method is used to return users login details template in CSV format.
func (uc *useCase) DownloadUserLoginDetailTemplate(ctx context.Context, fp string) (fileName, filePath string, err *error.Error) {

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

	dtlHeader := []string{"domain_name", "email_id", "employee_id", "isSingedup", "password", "role", "uuid"}
	csvWriter.Write(dtlHeader)
	return
}
