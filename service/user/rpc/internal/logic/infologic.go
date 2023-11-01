package logic

import (
	"context"

	"zouyi/thumbup/service/user/rpc/internal/svc"
	"zouyi/thumbup/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *user.InfoRequest) (*user.InfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.InfoResponse{}, nil
}
