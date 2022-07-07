package user_login_detail

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/entity"

	"github.com/jinzhu/copier"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// UploadUsersLoginDetail method is used to share user login detail to repository and verify the return data from repoistory.
func (uc *useCase) UploadUsersLoginDetail(ctx context.Context, file multipart.File) (savRes out.SaveResponse, err *error.Error) {
	var dtls []entity.UserLoginDetails
	details := []in.UploadUserDetail{}
	lineCnt := 0
	validate := validator.New()
	csvReader := csv.NewReader(file)
	for {
		detail, rerr := csvReader.Read()
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			err = error.NewInternal()
			return
		}
		lineCnt++
		if lineCnt != 1 {
			isSignedUp, cerr := strconv.Atoi(detail[3])
			if cerr != nil {
				err = error.NewBadRequest(fmt.Sprintf("Line no: %d, CSV file contain invalid input values", lineCnt))
				return
			}
			dtl := in.UploadUserDetail{
				DomainName:   detail[0],
				EmailID:      detail[1],
				EmployeeID:   detail[2],
				IsSignedUp:   isSignedUp,
				Password:     detail[4],
				Role:         detail[5],
				UUID:         detail[6],
				EnableAccess: "No",
			}

			if verr := validate.Struct(dtl); verr != nil {
				err = error.NewBadRequest(fmt.Sprintf("Line no: %d, Input validation was failed.", lineCnt))
				return
			}
			details = append(details, dtl)
		}
	}

	copier.Copy(&dtls, &details)

	rerr := uc.userLoginDetailRepo.CreateBulkUserLoginDetail(ctx, dtls)
	serr, ok := rerr.(*mysql.MySQLError)
	if rerr != nil && !ok {
		err = error.NewInternal()
		return
	}

	if ok && serr.Number == 1062 {
		err = error.NewBulkInsertConflict(serr.Message)
		return
	}

	savRes = out.SaveResponse{
		IsSaved: "true",
	}
	return
}
