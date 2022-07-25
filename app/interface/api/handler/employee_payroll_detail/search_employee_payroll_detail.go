package employee_payroll_detail

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

func (u *EmployeePayrollDetailHandler) SearchEmployeePayrollDetail(ctx *gin.Context) {

	filterMap := make(map[string]interface{})
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

	if employeeID, ok := ctx.GetQuery("employee_id"); ok {
		filterMap["employee_id"] = employeeID
	}
	if panNumber, ok := ctx.GetQuery("pan_number"); ok {
		filterMap["pan_number"] = panNumber
	}
	if uanNumberStr, ok := ctx.GetQuery("uan_number"); ok {
		uanNumber, cerr := strconv.Atoi(uanNumberStr)
		if cerr != nil {
			err := error.NewBadRequest("Invalid query pareameters. Verify request query parameters once.")
			log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
			ctx.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		}
		filterMap["uan_number"] = uanNumber
	}
	if bankAccountNumberStr, ok := ctx.GetQuery("bank_account_number"); ok {
		bankAccountNumber, cerr := strconv.Atoi(bankAccountNumberStr)
		if cerr != nil {
			err := error.NewBadRequest("Invalid query pareameters. Verify request query parameters once.")
			log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
			ctx.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		}
		filterMap["bank_account_number"] = bankAccountNumber
	}
	if bankIfscCode, ok := ctx.GetQuery("bank_ifsc_code"); ok {
		filterMap["bank_ifsc_code"] = bankIfscCode
	}
	if passportNumber, ok := ctx.GetQuery("passport_number"); ok {
		filterMap["passport_number"] = passportNumber
	}
	if pfAccountNumber, ok := ctx.GetQuery("pf_account_number"); ok {
		filterMap["pf_account_number"] = pfAccountNumber
	}
	if taxRegime, ok := ctx.GetQuery("tax_regime"); ok {
		filterMap["tax_regime"] = taxRegime
	}
	if effectiveFrom, ok := ctx.GetQuery("effective_from"); ok {
		filterMap["effective_from"] = effectiveFrom
	}
	if epsAccountNumber, ok := ctx.GetQuery("eps_account_number"); ok {
		filterMap["eps_account_number"] = epsAccountNumber
	}
	if prAccountNumber, ok := ctx.GetQuery("pr_account_number"); ok {
		filterMap["pr_account_number"] = prAccountNumber
	}
	if esiNumber, ok := ctx.GetQuery("esi_number"); ok {
		filterMap["esi_number"] = esiNumber
	}

	if len(filterMap) < 1 {
		err := error.NewBadRequest("Valid query pareameters not found. Verify request query parameters once.")
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	details, err := u.employeePayrollDetailUseCase.SearchEmployeePayrollDetail(ctx, filterMap)
	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, details)

}
