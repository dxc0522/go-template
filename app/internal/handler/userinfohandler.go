package handler

import (
	"net/http"

	"github.com/go-template/common/response"

	"github.com/go-template/app/internal/logic"
	"github.com/go-template/app/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx, r, &w)
		resp, err := l.UserInfo()
		response.Response(r, w, resp, err)
	}
}
