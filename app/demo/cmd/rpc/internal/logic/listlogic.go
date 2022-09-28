package logic

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-demo/app/demo/cmd/rpc/demo"
	"go-zero-demo/app/demo/cmd/rpc/internal/svc"
	"go-zero-demo/app/demo/cmd/rpc/pb"
	"go-zero-demo/app/demo/model"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

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
