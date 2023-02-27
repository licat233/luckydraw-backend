package winningRecords

import (
	"context"

	"luckydraw-backend/common/ctxdata"
	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLogic {
	return &QueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLogic) Query(req *types.QueryReq) (any, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	activityId := ctxdata.GetAidFromCtx(l.ctx)
	if activityId == 0 {
		activity, err := l.svcCtx.ActivityModel.FindOneByUuid(l.ctx, req.ActivityUuid)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询活动失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
		if err == model.ErrNotFound {
			activityId = 0
		} else {
			activityId = activity.Id
		}
	}
	if activityId == 0 {
		return respx.SingleResp("success", []*model.WinningRecords{}), nil
	}
	if userId == 0 && req.Passport != "" && activityId != 0 {
		user, err := l.svcCtx.UsersModel.FindsByPassportAndActivityId(l.ctx, req.Passport, activityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询用户失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
		if err == model.ErrNotFound {
			userId = 0
		} else {
			userId = user.Id
		}
	}
	if userId == 0 {
		//返回空
		return respx.SingleResp("success", []*model.WinningRecords{}), nil
	}

	list, err := l.QueryWinningRecords(userId, activityId)
	if err != nil {
		l.Logger.Errorf("查询中奖记录失败，错误信息：%v", err)
		return nil, errorx.InternalError(err)
	}
	var data = make([]*types.WinningRecordsInfo, 0)

	for _, v := range list {
		if v.AwardId != 0 {
			award, err := l.svcCtx.AwardsModel.FindOne(l.ctx, v.AwardId)
			if err != nil && err != model.ErrNotFound {
				l.Logger.Errorf("查询奖品失败，err:%v", err)
				return nil, errorx.InternalError(err)
			}
			if err == model.ErrNotFound {
				continue
			}
			pubAward := &types.PublicAwards{
				Id:    award.Id,
				Uuid:  award.Uuid,
				Grade: award.Grade,
				Name:  award.Name,
				Image: award.Image,
				Price: award.Price,
				IsWin: award.IsWin > 0,
			}
			Record := &types.WinningRecords{
				Id:         v.Id,
				UserId:     v.UserId,
				AwardId:    v.AwardId,
				ActivityId: v.ActivityId,
				Ip:         v.Ip,
				Platform:   v.Platform,
			}
			data = append(data, &types.WinningRecordsInfo{
				WinningRecord: Record,
				Award:         pubAward,
			})
		}
	}
	return respx.SingleResp("success", data), nil
}

func (l *QueryLogic) QueryWinningRecords(userId, activityId int64) (resp []*model.WinningRecords, err error) {
	resp, err = l.svcCtx.WinningRecordsModel.FindsByUserIdAndActivityId(l.ctx, userId, activityId)
	return
}
