package employee_payroll_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

func (u *EmployeePayrollDetailHandler) CreateEmployeePayrollDetail(ctx *gin.Context) {
	var payrolldetail in.CreateEmployeePayrollDetail

	if ok := handler.ValidateData(ctx, &payrolldetail); !ok {
		return
	}

	employeePayrollDetailRes, err := u.employeePayrollDetailUseCase.CreateEmployeePayrollDetail(ctx, payrolldetail)

	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, employeePayrollDetailRes)

}
