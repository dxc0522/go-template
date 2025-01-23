package logic

import (
	"context"
	"github.com/go-template/demo/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type DemoLogic struct {
	logx.Logger
	*svc.ServiceContext
	ctx        context.Context
	reqCtx     *http.Request
	respWriter *http.ResponseWriter
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext, reqCtx *http.Request, respWriter *http.ResponseWriter) *DemoLogic {
	return &DemoLogic{
		Logger:         logx.WithContext(ctx),
		ServiceContext: svcCtx,
		ctx:            ctx,
		reqCtx:         reqCtx,
		respWriter:     respWriter,
	}
}
