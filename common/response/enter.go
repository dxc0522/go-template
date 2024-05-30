package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response http返回
func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		//成功返回
		r := &Body{
			Code: http.StatusOK,
			Msg:  http.StatusText(http.StatusOK),
			Data: resp,
		}
		httpx.WriteJson(w, http.StatusOK, r)
		return
	}
	//错误返回
	httpx.WriteJson(w, http.StatusInternalServerError, &Body{
		Code: http.StatusInternalServerError,
		Msg:  err.Error(),
		Data: nil,
	})

}
