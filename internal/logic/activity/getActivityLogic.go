package activity

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

type GetActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityLogic {
	return &GetActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetActivityLogic) GetActivity(req *types.GetActivityReq) (any, error) {
	activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error("failed to find activity, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.New(tools.ActivityToResp(activity)), nil
}
