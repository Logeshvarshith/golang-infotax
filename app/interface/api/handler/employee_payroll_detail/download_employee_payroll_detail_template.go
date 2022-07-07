package employee_payroll_detail

import (
	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

type EmployeePayrollDetailHandler struct {
	employeePayrollDetailUseCase employee_payroll_detail.UseCaser
}

func NewEmployeePayrollDetailHandler(employeePayrollDetailUseCase employee_payroll_detail.UseCaser) EmployeePayrollDetailHandler {
	return EmployeePayrollDetailHandler{
		employeePayrollDetailUseCase: employeePayrollDetailUseCase,
	}
}

func (u *EmployeePayrollDetailHandler) DownloadEmployeePayrollDetailTemplate(ctx *gin.Context) {
	path := "outfile/employee_payroll_detail_template.csv"
	fileName, filePath, err := u.employeePayrollDetailUseCase.DownloadEmployeePayrollDetailTemplate(ctx, path)
	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "inline;filename="+fileName)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Cache-Control", "no-cache")
	ctx.File(filePath)

}
