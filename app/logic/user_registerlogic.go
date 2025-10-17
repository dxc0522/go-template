package logic

import (
	"github.com/go-template/app/dbmodels"
	"github.com/go-template/app/types"
)

func (l *AppLogic) Register(req *types.UserReq) error {
	return l.DB.Create(&dbmodels.Users{
		Username:     req.Name,
		Email:        req.Email,
		PasswordHash: req.Password,
	}).Error
}
