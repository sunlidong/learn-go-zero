package logic

import (
	"context"

	"jwttoken/internal/svc"
	"jwttoken/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EdituserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEdituserLogic(ctx context.Context, svcCtx *svc.ServiceContext) EdituserLogic {
	return EdituserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EdituserLogic) Edituser(req types.UserUpdateReq) (*types.UserOptResp, error) {
	// todo: add your logic here and delete this line

	return &types.UserOptResp{
		Id: req.Id,
	}, nil
}
