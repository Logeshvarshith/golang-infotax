package employee_payroll_detail

import (
	"net/http"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

func (u *EmployeePayrollDetailHandler) DeleteMultipleEmployeePayrollDetail(ctx *gin.Context) {

	var empID in.DeleteMultipleEmployee
	if ok := handler.ValidateData(ctx, &empID); !ok {
		return
	}

	delRes, err := u.employeePayrollDetailUseCase.DeleteMultipleEmployeePayrollDetail(ctx, empID)

	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, delRes)
}
