package test

import (
	"context"

	"go-zero-demo/app/demo/cmd/api/internal/svc"
	"go-zero-demo/app/demo/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdataLogic {
	return &UpdataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdataLogic) Updata(req *types.UpdateReq) (resp *types.UpdateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
