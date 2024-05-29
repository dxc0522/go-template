package logic

import (
	"context"
	"net/http"

	"github.com/go-template/app/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AppLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	reqCtx  *http.Request
	respCtx *http.ResponseWriter
}

func NewAppLogic(ctx context.Context, svcCtx *svc.ServiceContext, reqCtx *http.Request, respCtx *http.ResponseWriter) *AppLogic {
	return &AppLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		reqCtx:  reqCtx,
		respCtx: respCtx,
	}
}

func (l *AppLogic) App() (resp string, err error) {
	// todo: add your logic here and delete this line
	return "Hello World!", nil
}
