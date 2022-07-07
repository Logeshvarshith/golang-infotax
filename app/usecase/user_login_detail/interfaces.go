package user_login_detail

import (
	"context"
	"mime/multipart"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

type UseCaser interface {
	GetUserLoginDetail(ctx context.Context, empID string) (out.UserLoginDetail, *error.Error)
	GetUserLoginRole(ctx context.Context, empID string) (out.UserLoginRole, *error.Error)
	GetAllUserLoginDetail(ctx context.Context) ([]out.UserLoginDetail, *error.Error)
	DeleteUserLoginDetail(ctx context.Context, empID string) (out.DeleteResponse, *error.Error)
	CreateUserLoginDetail(ctx context.Context, detail in.CreateUserDetail) (out.SaveResponse, *error.Error)
	UpdateUserLoginDetail(ctx context.Context, empID string, detail in.UpdateUserDetail) (out.UpdateResponse, *error.Error)
	SearchUserLoginDetail(ctx context.Context, filterMap map[string]interface{}) ([]out.UserLoginDetail, *error.Error)
	DownloadUsersLoginDetail(ctx context.Context, filePath string) (string, string, *error.Error)
	DownloadUserLoginDetailTemplate(ctx context.Context, filePath string) (string, string, *error.Error)
	UploadUsersLoginDetail(ctx context.Context, file multipart.File) (out.SaveResponse, *error.Error)
}

type useCase struct {
	userLoginDetailRepo repository.UserLoginDetailRepository
}

// NewUseCase function is used to make new userCase struct.
func NewUseCase(userLoginDetailRepo repository.UserLoginDetailRepository) UseCaser {
	return &useCase{
		userLoginDetailRepo: userLoginDetailRepo,
	}
}
