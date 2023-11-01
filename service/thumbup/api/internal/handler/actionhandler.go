package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zouyi/thumbup/service/thumbup/api/internal/logic"
	"zouyi/thumbup/service/thumbup/api/internal/svc"
	"zouyi/thumbup/service/thumbup/api/internal/types"
)

func ActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewActionLogic(r.Context(), svcCtx)
		resp, err := l.Action(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
