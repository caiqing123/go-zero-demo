## 基本生成模块

### 1、goctl生成api代码

1）命令行进入app/demo/cmd/api/desc目录下。

```shell
$ goctl api go -api *.api -dir ../  -style=goZero # win 需指定.api文件
```

2）修改etc 配置文件

### 2、goctl生成rpc代码

1）命令行进入app/demo/cmd/rpc/pb目录下。

```shell
$ goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ # win 需指定.proto文件
```

2）修改etc 配置文件

### 3、model生成

1） 生成model代码，直接运行template/script/mysql/genModel.sh  参数


## api修改

1.添加api需要连接的rpc 修改 app\demo\cmd\api\internal\svc\serviceContext.go,对应Config也得修改

```go
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
```

2.调用rpc添加 修改 app\demo\cmd\api\internal\logic\test\listlogic.go 里逻辑 

```go
//list 例子 
func (l *ListLogic) List(req *types.ListReq) (*types.ListResp, error) {
    testResp, err := l.svcCtx.DemoRcp.List(l.ctx, &demo.ListReq{
        Mobile:   req.Mobile,
        Nickname: req.Nickname,
        Page:     req.Page,
        PageSize: req.PageSize,
    })
    
    if err != nil {
        return nil, err
    }
    
    var resp types.ListResp
    _ = copier.Copy(&resp, testResp)
    
    return &resp, nil
}
```

## rpc修改

1.添加rpc需要调用的数据库和模块 修改 app\demo\cmd\rpc\internal\svc\serviceContext.go,对应Config也得修改

```go
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

	DemoModel model.TestModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		DemoModel: model.NewTestModel(sqlConn, c.Cache),
	}
}
```

2.rpc处理逻辑 修改 app\demo\cmd\rpc\internal\logic\listlogic.go

```go
//list 例子 
func (l *ListLogic) List(in *pb.ListReq) (*pb.ListResp, error) {
    whereBuilder := l.svcCtx.TestModel.RowBuilder()
    where := map[string]interface{}{}
    if in.Mobile != "" {
        where["mobile"] = in.Mobile
    }
    if in.Nickname != "" {
        where["nickname"] = in.Nickname
    }
    whereBuilder = whereBuilder.Where(squirrel.Eq(where))
    
    test, err := l.svcCtx.TestModel.FindPageListByPage(l.ctx, whereBuilder, in.Page, in.PageSize, "")
    
    if err != nil && err != model.ErrNotFound {
        return nil, errors.Wrapf(err, "List find test db err , id:%d , err:%v", err.Error())
    }
    
    if test == nil {
        return &demo.ListResp{}, nil
    }
    
    count, _ := l.svcCtx.TestModel.FindCount(l.ctx, l.svcCtx.TestModel.CountBuilder("id").Where(squirrel.Eq(where)))
    
    var resp []*demo.Test
    
    _ = copier.Copy(&resp, test)
    
    return &demo.ListResp{
        TestInfo:   resp,
        TotalCount: count,
    }, nil
}
```