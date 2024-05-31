package handler

import (
	"net/http"

	"github.com/go-template/common/response"

	"github.com/go-template/app/internal/logic"
	"github.com/go-template/app/internal/svc"
	"github.com/go-template/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// App
func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx, r, &w)
		resp, err := l.UserInfo(&req)
		response.Response(r, w, resp, err)
	}
}
