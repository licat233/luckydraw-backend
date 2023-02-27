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

type GetActivityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityListLogic {
	return &GetActivityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetActivityListLogic) GetActivityList(req *types.GetActivityListReq) (any, error) {
	in := &model.Activity{
		Id:     req.Id,
		Uuid:   req.Uuid,
		Name:   req.Name,
		Status: req.Status,
	}
	list, total, err := l.svcCtx.ActivityModel.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
	if err != nil {
		l.Logger.Error("获取活动列表失败, err: ", err)
		return nil, errorx.InternalError(err)
	}
	data := make([]*types.Activity, 0)
	for _, v := range list {
		data = append(data, tools.ActivityToResp(v))
	}
	return respx.NewListData(data, total, req.PageSize, req.Page), nil
}
