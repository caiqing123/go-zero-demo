package test

import (
	"context"

	"github.com/jinzhu/copier"

	"go-zero-demo/app/demo/cmd/api/internal/svc"
	"go-zero-demo/app/demo/cmd/api/internal/types"
	"go-zero-demo/app/demo/cmd/rpc/demo"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.DelReq) (*types.DelResp, error) {
	testResp, err := l.svcCtx.DemoRcp.Del(l.ctx, &demo.DelReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	var resp types.DelResp
	_ = copier.Copy(&resp, testResp)

	return &resp, nil
}
