package user_login_detail

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

// UploadUsersLoginDetail godoc
// @Summary Insert bulk of users login details present in input csv file
// @Description Insert bulk of users login details present in input csv file
// @Tags User Login Detail
// @ID upload_users_login_detail
// @Accept application/octet-stream
// @Produce application/json
// @Param file formData file true "Users login details form"
// @Success 200 {object} out.SaveResponse
// @Failure 404 {object} error.Error
// @Failure 409 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /infotax/user_login_detail/upload [post]
func (u *UserLoginDetailHandler) UploadUsersLoginDetail(ctx *gin.Context) {

	file, fileHeader, ferr := ctx.Request.FormFile("file")
	if ferr != nil {
		err := error.NewBadRequest("Invalid request file. Verify request file.")
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ext := filepath.Ext(fileHeader.Filename)
	if ext != ".csv" {
		err := error.NewBadRequest("Invalid request file extension. Verify request file extension.")
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	defer file.Close()

	saveRes, err := u.userLoginDetailUseCase.UploadUsersLoginDetail(ctx, file)
	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, saveRes)
}
