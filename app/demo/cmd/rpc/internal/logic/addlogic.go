package logic

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"go-zero-demo/app/demo/cmd/rpc/demo"
	"go-zero-demo/app/demo/cmd/rpc/internal/svc"
	"go-zero-demo/app/demo/cmd/rpc/pb"
	"go-zero-demo/app/demo/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *pb.AddReq) (*pb.AddResp, error) {
	test, err := l.svcCtx.TestModel.FindOneByQuery(l.ctx, l.svcCtx.TestModel.RowBuilder().Where(squirrel.Eq{"mobile": in.Mobile}))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(err, "mobile:%s,err:%v", in.Mobile, err)
	}
	if test != nil {
		return nil, errors.Wrapf(errors.New("手机号存在"), "add test exists mobile:%s,err:%v", in.Mobile, err)
	}

	var userId int64
	if err := l.svcCtx.TestModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		test := new(model.Test)
		test.Mobile = in.Mobile
		test.Nickname = in.Nickname

		insertResult, err := l.svcCtx.TestModel.Insert(ctx, session, test)
		if err != nil {
			return errors.Wrapf(err, "add db test Insert err:%v,user:%+v", err, test)
		}
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			return errors.Wrapf(err, "add db test insertResult.LastInsertId err:%v,user:%+v", err, test)
		}
		userId = lastId

		return nil
	}); err != nil {
		return nil, err
	}

	var resp demo.Test

	resp.Mobile = in.Mobile
	resp.Nickname = in.Nickname
	resp.Id = userId

	return &pb.AddResp{
		TestInfo: &resp,
	}, nil
}
