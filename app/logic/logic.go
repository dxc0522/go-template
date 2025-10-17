package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-template/app/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type AppLogic struct {
	logx.Logger
	*svc.ServiceContext
	ctx        context.Context
	reqCtx     *http.Request
	respWriter *http.ResponseWriter
}

func NewAppLogic(ctx context.Context, svcCtx *svc.ServiceContext, reqCtx *http.Request, respWriter *http.ResponseWriter) *AppLogic {
	return &AppLogic{
		Logger:         logx.WithContext(ctx),
		ServiceContext: svcCtx,
		ctx:            ctx,
		reqCtx:         reqCtx,
		respWriter:     respWriter,
	}
}

type ResponseDataError struct {
	Data any `json:"data"`
	Code int `json:"code"`
}

func (e *ResponseDataError) Error() string {
	jsonStr, err := json.Marshal(e.Data)
	if err != nil {
		return fmt.Sprint(e.Data)
	}
	return string(jsonStr)
}

func NewErrorResponse(code int, resp any) *ResponseDataError {
	return &ResponseDataError{
		Data: resp,
		Code: code,
	}
}
