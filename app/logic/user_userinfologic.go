package logic

import (
	"errors"

	"gorm.io/gorm"

	"github.com/go-template/app/dbmodels"
	"github.com/go-template/app/types"
)

func (l *AppLogic) UserInfo(req *types.UserInfoReq) (res *types.UserResp, err error) {
	user := dbmodels.Users{}
	err = l.DB.Where("id=?", req.Id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return
	}
	return &types.UserResp{
		Name:  user.Username,
		Email: user.Email,
	}, nil
}
