package user

import (
	"net/http"

	"MyRainFarm/application/user/api/internal/logic/user"
	"MyRainFarm/application/user/api/internal/svc"
	"MyRainFarm/application/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RepwdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResetPwdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewRepwdLogic(r.Context(), svcCtx)
		resp, err := l.Repwd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
