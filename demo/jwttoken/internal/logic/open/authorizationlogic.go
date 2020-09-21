package logic

import (
	"context"
	"jwttoken/internal/utils"
	"time"

	"jwttoken/internal/svc"
	"jwttoken/internal/types"

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
	token, err := utils.EncodeJwtToken(map[string]interface {
	}{
		"id": req.Mobile,
	})
	return &types.UserOptResp{
		Token: token,
		Id:    uint(time.Now().Unix()),
	}, err
}
