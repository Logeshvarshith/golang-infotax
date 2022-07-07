package user_login_detail

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

// SearchUserLoginDetail godoc
// @Summary Search login user detail based on search filter
// @Description Search login user detail based on search filter
// @Tags User Login Detail
// @ID search_user_login_detail
// @Accept json
// @Produce json
// @Param employee_id query string false "EmployeeID filter"
// @Param domain_name query string false "DomainName filter"
// @Param email_id query string false "EmailID filter"
// @Param password query string false "Password filter"
// @Param uuid query string false "UUID filter"
// @Param isSignedup query int false "IsSignedup filter"
// @Param role query string false "Role filter"
// @Success 200 {object} []out.UserLoginDetail
// @Failure 400 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /infotax/user_login_detail/search [get]
func (u *UserLoginDetailHandler) SearchUserLoginDetail(ctx *gin.Context) {

	filterMap := make(map[string]interface{})

	if employeeID, ok := ctx.GetQuery("employee_id"); ok {
		filterMap["employee_id"] = employeeID
	}
	if domainName, ok := ctx.GetQuery("domain_name"); ok {
		filterMap["domain_name"] = domainName
	}
	if emailID, ok := ctx.GetQuery("email_id"); ok {
		filterMap["email_id"] = emailID
	}
	if password, ok := ctx.GetQuery("password"); ok {
		filterMap["password"] = password
	}
	if uuid, ok := ctx.GetQuery("uuid"); ok {
		filterMap["uuid"] = uuid
	}
	if isSignedUpStr, ok := ctx.GetQuery("isSignedup"); ok {
		isSignedUp, cerr := strconv.Atoi(isSignedUpStr)
		if cerr != nil {
			err := error.NewBadRequest("Invalid query pareameters. Verify request query parameters once.")
			log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
			ctx.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		}
		filterMap["isSignedup"] = isSignedUp
	}
	if role, ok := ctx.GetQuery("role"); ok {
		filterMap["role"] = role
	}

	if len(filterMap) < 1 {
		err := error.NewBadRequest("Valid query pareameters not found. Verify request query parameters once.")
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	details, err := u.userLoginDetailUseCase.SearchUserLoginDetail(ctx, filterMap)
	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, details)

}
