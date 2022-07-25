package employee_payroll_detail

import (
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"

	"github.com/gin-gonic/gin"
)

func (u *EmployeePayrollDetailHandler) DownloadEmployeePayrollDetail(ctx *gin.Context) {
	path := "outfile/employee_payroll_detail.csv"
	fileName, filePath, err := u.employeePayrollDetailUseCase.DownloadEmployeePayrollDetail(ctx, path)
	if err != nil {
		log.Logger.Errorf("Type : %v, Message : %v\n", err.Type, err.Message)
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
