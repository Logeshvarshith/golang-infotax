package employee_payroll_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

func (u *EmployeePayrollDetailHandler) DeleteEmployeeOfficialDetail(ctx *gin.Context) {
	empID := ctx.Param("emp_id")

	delRes, err := u.employeePayrollDetailUseCase.DeleteEmployeePayrollDetail(ctx, empID)

	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, delRes)
}
