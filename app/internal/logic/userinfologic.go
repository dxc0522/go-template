package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-template/app/internal/svc"
	"github.com/go-template/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	reqCtx  *http.Request
	respCtx *http.ResponseWriter
}

// App
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext, reqCtx *http.Request, respCtx *http.ResponseWriter) *UserInfoLogic {
	return &UserInfoLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		reqCtx:  reqCtx,
		respCtx: respCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	l.Info(req)
	userId := l.ctx.Value("user_id").(json.Number)
	fmt.Printf("%v %T", userId, userId)
	userName := l.ctx.Value("username").(string)
	fmt.Printf("%v %T", userName, userName)
	uid, _ := userId.Int64()
	return &types.UserInfoResponse{
		Id:       uint(uid),
		UserName: userName,
	}, nil
}
