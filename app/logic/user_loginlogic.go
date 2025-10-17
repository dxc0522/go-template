package logic

import (
	"errors"

	"gorm.io/gorm"

	"github.com/go-template/app/dbmodels"
	"github.com/go-template/app/types"
)

func (l *AppLogic) Login(req *types.LoginReq) (res *types.LoginResp, err error) {
	user := dbmodels.Users{}
	err = l.DB.Where("email=? and password_hash=?", req.Email, req.Password).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	l.Info("user", user)
	return &types.LoginResp{
		Token: "123",
	}, nil
}
