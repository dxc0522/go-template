package logic

import (
	"context"
	"errors"
	"github.com/go-template/app/dbmodel"
	"github.com/go-template/common/jwts"
	"gorm.io/gorm"
	"net/http"

	"github.com/go-template/app/internal/svc"
	"github.com/go-template/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	reqCtx *http.Request
	respW  *http.ResponseWriter
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, reqCtx *http.Request, respW *http.ResponseWriter) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		reqCtx: reqCtx,
		respW:  respW,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// check user
	exitUser := dbmodel.Users{}
	err = l.svcCtx.DB.Where(&dbmodel.Users{
		Name: req.UserName,
	}).Take(&exitUser).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil || exitUser.Id > 0 {
		return "exit user name", errors.New("exit user name")
	}
	// save user to db
	user := dbmodel.Users{
		Name:     req.UserName,
		Password: req.Password,
		Mobile:   req.Mobile,
	}
	err = l.svcCtx.DB.Save(&user).Error
	if err != nil {
		return "error", err
	}
	// set token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		Username: req.UserName,
		UserID:   uint(user.Id),
		Role:     1,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return "error", err
	}
	w := *l.respW
	w.Header().Set("Authorization", "Bearer "+token)
	return "success", nil
}
