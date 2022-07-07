package framework

import (
	"context"
	"time"

	employee_official_detail "www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler/employee_official_detail"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler/user_login_detail"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler/middleware"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/config"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/registry"
)

// Handler is used to initialize all the REST endpoints.
func Handler(conf *config.Config) *gin.Engine {

	r := gin.Default()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r.Use(middleware.CORS())

	// Swagger Endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userLoginDetailGroup := r.Group(conf.UserLoginDetailBaseUrl)
	userLoginDetailHandler := user_login_detail.NewUserLoginDetailHandler(
		registry.InjectedUserLoginDetailUseCase(ctx),
	)

	userLoginDetailGroup.GET("/:emp_id", userLoginDetailHandler.GetUserLoginDetail)
	userLoginDetailGroup.GET("/role/:emp_id", userLoginDetailHandler.GetUserLoginRole)
	userLoginDetailGroup.GET("/search", userLoginDetailHandler.SearchUserLoginDetail)
	userLoginDetailGroup.GET("/", userLoginDetailHandler.GetAllUserLoginDetail)
	userLoginDetailGroup.DELETE("/:emp_id", userLoginDetailHandler.DeleteUserLoginDetail)
	userLoginDetailGroup.POST("/", userLoginDetailHandler.CreateUserLoginDetail)
	userLoginDetailGroup.PUT("/:emp_id", userLoginDetailHandler.UpdateUserLoginDetail)
	userLoginDetailGroup.GET("/download", userLoginDetailHandler.DownloadUsersLoginDetail)
	userLoginDetailGroup.GET("/download/template", userLoginDetailHandler.DownloadUserLoginDetailTemplate)
	userLoginDetailGroup.POST("/upload", userLoginDetailHandler.UploadUsersLoginDetail)

	EmployeeOfficialDetailGroup := r.Group(conf.EmployeeOfficialDetailBaseUrl)
	EmployeeOfficialDetailHandler := employee_official_detail.NewEmployeeofficialDetailHandler(
		registry.InjectedEmployeeOfficialDetailUseCase(ctx),
	)

	EmployeeOfficialDetailGroup.POST("/", EmployeeOfficialDetailHandler.CreateEmployeeOfficialDetail)

	return r

}
