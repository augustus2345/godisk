package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_disk/core/internal/logic"
	"go_disk/core/internal/svc"
	"go_disk/core/internal/types"
)

func UserDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserDetailsLogic(r.Context(), svcCtx)
		resp, err := l.UserDetails(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
