package adminer

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAdminerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAdminerLogic {
	return &DelAdminerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAdminerLogic) DelAdminer(req *types.DelAdminerReq) (any, error) {
	if err := l.svcCtx.AdminerModel.Delete(l.ctx, req.Id); err != nil {
		l.Logger.Errorf("删除管理员失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
