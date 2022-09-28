package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"go-zero-demo/app/demo/cmd/api/internal/config"
	"go-zero-demo/app/demo/cmd/rpc/demo"
)

type ServiceContext struct {
	Config  config.Config
	DemoRcp demo.Demo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DemoRcp: demo.NewDemo(zrpc.MustNewClient(c.DemoRpcConf)),
	}
}
