package adminer

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/logic/tools"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerLogic {
	return &GetAdminerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminerLogic) GetAdminer(req *types.GetAdminerReq) (any, error) {
	adminer, err := l.svcCtx.AdminerModel.FindOne(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error("failed to find adminer, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.New(tools.AdminerToResp(adminer)), nil
}
