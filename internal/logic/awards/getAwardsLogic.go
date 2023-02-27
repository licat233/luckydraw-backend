package awards

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

type GetAwardsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAwardsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAwardsLogic {
	return &GetAwardsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAwardsLogic) GetAwards(req *types.GetAwardsReq) (any, error) {
	award, err := l.svcCtx.AwardsModel.FindOne(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error("failed to find award, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.New(tools.AwardsToResp(award)), nil
}
