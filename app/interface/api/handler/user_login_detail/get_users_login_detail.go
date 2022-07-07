package user_login_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

// GetAllUserLoginDetail godoc
// @Summary Get all login user detail
// @Description Get all login user detail
// @Tags User Login Detail
// @ID get_all_user_login_detail
// @Accept json
// @Produce json
// @Success 200 {object} []out.UserLoginDetail
// @Failure 500 {object} error.Error
// @Router /infotax/user_login_detail/ [get]
func (u *UserLoginDetailHandler) GetAllUserLoginDetail(ctx *gin.Context) {

	userLoginDtls, err := u.userLoginDetailUseCase.GetAllUserLoginDetail(ctx)
	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, userLoginDtls)

}
