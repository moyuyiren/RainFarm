package indexContent

import (
	"net/http"

	"MyRainFarm/application/indexContent/api/internal/logic/indexContent"
	"MyRainFarm/application/indexContent/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func IndexContentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := indexContent.NewIndexContentLogic(r.Context(), svcCtx)
		resp, err := l.IndexContent()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
