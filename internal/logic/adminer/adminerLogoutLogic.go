package adminer

import (
	"context"

	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminerLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminerLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminerLogoutLogic {
	return &AdminerLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminerLogoutLogic) AdminerLogout() (resp *types.BaseResp, err error) {
	return
}
