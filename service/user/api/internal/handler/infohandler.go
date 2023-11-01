package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zouyi/thumbup/service/user/api/internal/logic"
	"zouyi/thumbup/service/user/api/internal/svc"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
