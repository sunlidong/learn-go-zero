package logic

import (
	"context"

	"hello/internal/svc"
	"hello/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) AuthorizationLogic {
	return AuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthorizationLogic) Authorization(req types.UserOptReq) (*types.UserOptResp, error) {
	// todo: add your logic here and delete this line

	return &types.UserOptResp{}, nil
}
