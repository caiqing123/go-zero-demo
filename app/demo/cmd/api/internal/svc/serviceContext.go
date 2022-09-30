package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"go-zero-demo/app/demo/cmd/api/internal/config"
	"go-zero-demo/app/demo/cmd/api/internal/middleware"
	"go-zero-demo/app/demo/cmd/rpc/demo"
)

type ServiceContext struct {
	Config  config.Config
	DemoRcp demo.Demo
	ForceUa rest.Middleware
	DBModel sqlc.CachedConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:  c,
		ForceUa: middleware.NewForceUaMiddleware().Handle,
		DemoRcp: demo.NewDemo(zrpc.MustNewClient(c.DemoRpcConf)),
		DBModel: sqlc.NewConn(sqlConn, c.Cache),
	}
}
