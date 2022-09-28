package logic

import (
	"context"

	"github.com/pkg/errors"

	"go-zero-demo/app/demo/cmd/rpc/internal/svc"
	"go-zero-demo/app/demo/cmd/rpc/pb"
	"go-zero-demo/app/demo/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLogic) Del(in *pb.DelReq) (*pb.DelResp, error) {
	test, err := l.svcCtx.TestModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(err, "del find test db err , id:%d , err:%v", in.Id, err)
	}
	if test == nil {
		return nil, errors.Wrapf(errors.New("不存在"), "id:%d", in.Id)
	}

	err = l.svcCtx.TestModel.Delete(l.ctx, test.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to del test db err:%v , in : %v", err, in)
	}

	return &pb.DelResp{
		Msg: "删除成功",
	}, nil
}
