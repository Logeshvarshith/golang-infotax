package employee_payroll_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

func (u *EmployeePayrollDetailHandler) GetAllEmployeePayrollDetail(ctx *gin.Context) {
	emppayrolldetail, err := u.employeePayrollDetailUseCase.GetAllEmployeePayrollDetail(ctx)

	if err != nil {
		log.Logger.Errorf("Type:%v,Message:%v", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, emppayrolldetail)
}
