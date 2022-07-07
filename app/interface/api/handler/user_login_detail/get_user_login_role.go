package user_login_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

// GetUserLoginRole godoc
// @Summary Get login user role based on employee id
// @Description Get role by employee id
// @Tags User Login Detail
// @ID get_user_login_detail_role
// @Accept json
// @Produce json
// @Param emp_id path string true "Employee ID"
// @Success 200 {object} out.UserLoginRole
// @Failure 400 {object} error.Error
// @Failure 404 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /infotax/user_login_detail/role/{emp_id} [get]
func (u *UserLoginDetailHandler) GetUserLoginRole(ctx *gin.Context) {

	empID := ctx.Param("emp_id")
	role, err := u.userLoginDetailUseCase.GetUserLoginRole(ctx, empID)
	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, role)

}
