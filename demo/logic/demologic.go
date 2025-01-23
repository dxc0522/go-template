package logic

import (
	"github.com/go-template/common/pkg/typex"
	"net/http"

	"github.com/go-template/common/pkg/errorx"
	"github.com/go-template/demo/types"
)

func (l *DemoLogic) Demo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	l.Info("log info", "name", "info")
	l.Error("log error", "name", "error")

	return nil, errorx.WithCodeResponse(http.StatusInternalServerError, typex.ErrorBody{
		Code: 1234,
		Msg:  "test custom error",
	})
}
