package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"go-zero-demo/app/demo/cmd/rpc/demo"
	"go-zero-demo/app/demo/cmd/rpc/internal/svc"
	"go-zero-demo/app/demo/cmd/rpc/pb"
	"go-zero-demo/app/demo/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *pb.UpdateReq) (*pb.UpdateResp, error) {
	test, err := l.svcCtx.TestModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(err, "Update find test db err , id:%d , err:%v", in.Id, err)
	}
	if test == nil {
		return nil, errors.Wrapf(errors.New("不存在"), "id:%d", in.Id)
	}

	test.Nickname = in.Nickname
	test.Mobile = in.Mobile

	err = l.svcCtx.TestModel.Update(l.ctx, test)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to update test db err:%v , in : %v", err, in)
	}

	var resp demo.Test
	_ = copier.Copy(&resp, test)

	return &pb.UpdateResp{
		TestInfo: &resp,
	}, nil
}
