package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"go-zero-demo/common/errorx"
)

type ForceUaMiddleware struct {
}

func NewForceUaMiddleware() *ForceUaMiddleware {
	return &ForceUaMiddleware{}
}

func (m *ForceUaMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取 User-Agent 标头信息
		if len(r.Header.Get("User-Agent")) == 0 {
			w.WriteHeader(http.StatusNotFound)
			httpx.Error(w, errorx.NewDefaultError("User-Agent 标头未找到"))
			return
		}

		next(w, r)
	}
}
