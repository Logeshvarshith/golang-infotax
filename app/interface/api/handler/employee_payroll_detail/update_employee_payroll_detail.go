package employee_payroll_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

func (u *EmployeePayrollDetailHandler) UpdateEmployeePayrollDetail(ctx *gin.Context) {
	empid := ctx.Param("emp_id")
	var detail in.UpdatedEmployeePayrollDetail
	if ok := handler.ValidateData(ctx, &detail); !ok {
		return
	}
	updRes, err := u.employeePayrollDetailUseCase.UpdateEmployeePayrollDetail(ctx, empid, detail)

	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, updRes)
}
