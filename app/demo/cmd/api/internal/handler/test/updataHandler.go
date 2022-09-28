package test

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/app/demo/cmd/api/internal/logic/test"
	"go-zero-demo/app/demo/cmd/api/internal/svc"
	"go-zero-demo/app/demo/cmd/api/internal/types"
)

func UpdataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := test.NewUpdataLogic(r.Context(), svcCtx)
		resp, err := l.Updata(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
