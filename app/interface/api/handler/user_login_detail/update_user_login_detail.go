package user_login_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

// UpdateUserLoginDetail godoc
// @Summary Update login user details
// @Description Update login user details
// @Tags User Login Detail
// @ID update_user_login_detail
// @Accept json
// @Produce json
// @Param emp_id path string true "Employee ID"
// @Param message body in.UpdateUserDetail true "Update user login details"
// @Success 200 {object} out.UpdateResponse
// @Failure 400 {object} error.Error
// @Failure 404 {object} error.Error
// @Failure 415 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /infotax/user_login_detail/{emp_id} [put]
func (u *UserLoginDetailHandler) UpdateUserLoginDetail(ctx *gin.Context) {

	empID := ctx.Param("emp_id")
	var detail in.UpdateUserDetail
	if ok := handler.ValidateData(ctx, &detail); !ok {
		return
	}

	updRes, err := u.userLoginDetailUseCase.UpdateUserLoginDetail(ctx, empID, detail)
	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, updRes)

}
