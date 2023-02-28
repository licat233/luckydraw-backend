package winningRecords

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

type GetWinningRecordsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWinningRecordsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWinningRecordsListLogic {
	return &GetWinningRecordsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWinningRecordsListLogic) GetWinningRecordsList(req *types.GetWinningRecordsListReq) (any, error) {
	in := &model.WinningRecords{
		Id:         req.Id,
		UserId:     req.UserId,
		AwardId:    req.AwardId,
		ActivityId: req.ActivityId,
		Ip:         req.Ip,
		Platform:   req.Platform,
	}
	list, total, err := l.svcCtx.WinningRecordsModel.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
	if err != nil {
		l.Logger.Errorf("获取中奖记录列表失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	var data []*types.WinningRecordsDetail
	for _, record := range list {
		award, err := l.svcCtx.AwardsModel.FindOne(l.ctx, record.AwardId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询奖品失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
		if err == model.ErrNotFound {
			award = nil
		}
		user, err := l.svcCtx.UsersModel.FindOne(l.ctx, record.UserId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询用户失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
		if err == model.ErrNotFound {
			user = nil
		}
		activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, record.ActivityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询活动失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
		if err == model.ErrNotFound {
			activity = nil
		}
		awd := tools.AwardsToResp(award)
		if awd == nil {
			continue
		}
		act := tools.ActivityToResp(activity)
		if act == nil {
			continue
		}
		usr := tools.UserToResp(user)
		if usr == nil {
			continue
		}
		data = append(data, &types.WinningRecordsDetail{
			WinningRecord: tools.WinningRecordsToResp(record),
			Award:         awd,
			User:          usr,
			Activity:      act,
		})
	}
	return respx.NewListData(data, total, req.PageSize, req.Page), nil
}
