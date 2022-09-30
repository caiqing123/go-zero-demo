package test

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"

	"go-zero-demo/app/demo/cmd/api/internal/svc"
	"go-zero-demo/app/demo/cmd/api/internal/types"
	"go-zero-demo/app/demo/cmd/rpc/demo"
	"go-zero-demo/common/validator"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateReq) (*types.UpdateResp, error) {
	//验证
	validator.DBModel = l.svcCtx.DBModel
	result := validator.NewValidator().Validate(req, "zh")
	if result != "" {
		return nil, fmt.Errorf(result)
	}

	testResp, err := l.svcCtx.DemoRcp.Update(l.ctx, &demo.UpdateReq{
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Id:       req.Id,
	})

	if err != nil {
		return nil, err
	}

	var resp types.UpdateResp
	_ = copier.Copy(&resp, testResp)

	return &resp, nil
}
