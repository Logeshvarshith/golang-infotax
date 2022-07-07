package user_login_detail

import (
	"net/http"

	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
)

// CreateUserLoginDetail godoc
// @Summary Create login user details
// @Description Create login user details
// @Tags User Login Detail
// @ID create_user_login_detail
// @Accept json
// @Produce json
// @Param message body in.CreateUserDetail true "User login details"
// @Success 200 {object} out.SaveResponse
// @Failure 400 {object} error.Error
// @Failure 409 {object} error.Error
// @Failure 415 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /infotax/user_login_detail [post]
func (u *UserLoginDetailHandler) CreateUserLoginDetail(ctx *gin.Context) {

	var detail in.CreateUserDetail
	if ok := handler.ValidateData(ctx, &detail); !ok {
		return
	}

	// Default values
	detail.Password = ""
	detail.IsSignedUp = 0
	detail.UUID = "TestData"
	detail.EnableAccess = "No"

	saveRes, err := u.userLoginDetailUseCase.CreateUserLoginDetail(ctx, detail)
	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, saveRes)

}
