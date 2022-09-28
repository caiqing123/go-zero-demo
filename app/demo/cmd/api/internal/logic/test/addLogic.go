package test

import (
	"context"

	"github.com/jinzhu/copier"

	"go-zero-demo/app/demo/cmd/api/internal/svc"
	"go-zero-demo/app/demo/cmd/api/internal/types"
	"go-zero-demo/app/demo/cmd/rpc/demo"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (*types.AddResp, error) {
	testResp, err := l.svcCtx.DemoRcp.Add(l.ctx, &demo.AddReq{
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
	})

	if err != nil {
		return nil, err
	}

	var resp types.AddResp
	_ = copier.Copy(&resp, testResp)

	return &resp, nil
}
