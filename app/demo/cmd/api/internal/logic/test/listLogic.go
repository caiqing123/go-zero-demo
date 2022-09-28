package test

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-demo/app/demo/cmd/api/internal/svc"
	"go-zero-demo/app/demo/cmd/api/internal/types"
	"go-zero-demo/app/demo/cmd/rpc/demo"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

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
