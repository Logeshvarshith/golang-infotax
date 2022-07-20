package employee_official_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"
)

func (u *EmployeeOfficialDetailHandler) UpdateEmployeeOfficialDetail(ctx *gin.Context) {
	empid := ctx.Param("emp_id")
	var detail in.CreateEmployeeOfficial
	if ok := handler.ValidateData(ctx, &detail); !ok {
		return
	}
	updRes, err := u.employeeofficialDetailUseCase.UpdateEmployeeOfficialDetail(ctx, empid, detail)

	if err != nil {
		log.Logger.Errorf("Type: %v, Message : %v\n", err.Type, err.Message)
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, updRes)
}
