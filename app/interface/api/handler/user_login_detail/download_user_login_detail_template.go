package user_login_detail

import (
	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

// DownloadUserLoginDetailTemplate godoc
// @Summary Return user login details template in csv file format
// @Description Return user login details template in csv file format
// @Tags User Login Detail
// @ID download_users_login_detail_template
// @Produce application/octet-stream
// @Success 200 {file} file
// @Failure 404 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /infotax/user_login_detail/download/template [get]
func (u *UserLoginDetailHandler) DownloadUserLoginDetailTemplate(ctx *gin.Context) {

	path := "outfile/user_login_detail_template.csv"
	fileName, filePath, err := u.userLoginDetailUseCase.DownloadUserLoginDetailTemplate(ctx, path)
	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "inline;filename="+fileName)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Cache-Control", "no-cache")
	ctx.File(filePath)
}
