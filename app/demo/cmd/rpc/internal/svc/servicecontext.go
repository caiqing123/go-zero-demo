package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"go-zero-demo/app/demo/cmd/rpc/internal/config"
	"go-zero-demo/app/demo/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	TestModel model.TestModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		TestModel: model.NewTestModel(sqlConn, c.Cache),
	}
}
