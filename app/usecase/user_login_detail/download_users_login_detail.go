package user_login_detail

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"path"

	"github.com/jinzhu/copier"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// DownloadUsersLoginDetail method is used to return users login details in CSV format.
func (uc *useCase) DownloadUsersLoginDetail(ctx context.Context, fp string) (fileName, filePath string, err *error.Error) {
	var details []out.UserLoginDetail
	dtls, rerr := uc.userLoginDetailRepo.GetAllUserLoginDetail(ctx)
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

	dtlHeader := []string{"domain_name", "email_id", "employee_id", "isSingedup", "password", "role", "uuid"}
	csvWriter.Write(dtlHeader)

	for _, detail := range details {
		dtl := []string{detail.DomainName, detail.EmailID, detail.EmployeeID, fmt.Sprintf("%v", detail.IsSignedUp), detail.Password, detail.Role, detail.UUID}
		csvWriter.Write(dtl)
	}

	return
}
