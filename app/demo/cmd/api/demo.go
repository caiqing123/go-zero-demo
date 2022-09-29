package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"

	"go-zero-demo/common/errorx"

	"go-zero-demo/app/demo/cmd/api/internal/config"
	"go-zero-demo/app/demo/cmd/api/internal/handler"
	"go-zero-demo/app/demo/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/demo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		causeErr := errors.Cause(err)                  // err类型
		if e, ok := causeErr.(*errorx.CodeError); ok { //自定义错误类型
			//自定义CodeError
			return http.StatusOK, e.Data()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := int(gstatus.Code())
				return http.StatusOK, errorx.CodeError{Code: grpcCode, Msg: gstatus.Message()}
			}
		}
		return http.StatusOK, errorx.CodeError{Code: errorx.DefaultCode, Msg: err.Error()}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
