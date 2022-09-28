package logic

import (
	"context"

	"go-zero-demo/app/demo/cmd/rpc/internal/svc"
	"go-zero-demo/app/demo/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdataLogic {
	return &UpdataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdataLogic) Updata(in *pb.UpdateReq) (*pb.UpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateResp{}, nil
}
