package logic

import (
	"context"

	"file/internal/svc"
	"file/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AuthorizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) AuthorizationLogic {
	return AuthorizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthorizationLogic) Authorization(req types.UserOptReq) (*types.UserOptResp, error) {
	// todo: add your logic here and delete this line

	return &types.UserOptResp{}, nil
}
