package logic

import (
	"context"
	"github.com/go-template/common/jwts"
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
	// todo: add your logic here and delete this line
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		Username: req.UserName,
		UserID:   1,
		Role:     1,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return "error", err
	}
	w := *l.respW
	w.Header().Set("Authorization", "Bearer "+token)
	return "success", nil
}
