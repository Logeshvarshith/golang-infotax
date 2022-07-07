package user_login_detail

import (
	"context"

	"github.com/jinzhu/copier"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// SearchUserLoginDetail method is used to send filter detail to repository and get login user details
func (uc *useCase) SearchUserLoginDetail(ctx context.Context, filterMap map[string]interface{}) (details []out.UserLoginDetail, err *error.Error) {
	dtls, rerr := uc.userLoginDetailRepo.SearchUserLoginDetail(ctx, filterMap)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	copier.Copy(&details, &dtls)

	return
}
