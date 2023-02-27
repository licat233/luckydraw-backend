/*
 * @Author: licat
 * @Date: 2023-02-21 11:12:10
 * @LastEditors: licat
 * @LastEditTime: 2023-02-21 12:13:13
 * @Description: licat233@gmail.com
 */
package awards

import (
	"context"
	"errors"
	"strings"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAwardsListByActivityUuidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAwardsListByActivityUuidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAwardsListByActivityUuidLogic {
	return &GetAwardsListByActivityUuidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAwardsListByActivityUuidLogic) GetAwardsListByActivityUuid(req *types.GetAwardsListByActivityUuidReq) (any, error) {
	activityUuid := strings.TrimSpace(req.ActivityUuid)
	if activityUuid == "" {
		return nil, errorx.ExternalError(types.ErrParams)
	}
	activity, err := l.svcCtx.ActivityModel.FindOneByUuid(l.ctx, activityUuid)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Errorf("查询活动失败，err:%v", err)
		return nil, errorx.InternalError(err)
	}
	//如果活动不存在
	if err == model.ErrNotFound {
		return nil, errorx.ExternalError(errors.New("the activity not found"))
	}
	activityId := activity.Id
	awards, err := l.svcCtx.AwardsModel.FindsByActivityId(l.ctx, activityId)
	if err != nil {
		l.Logger.Errorf("查询奖品列表失败，err:%v", err)
		return nil, errorx.InternalError(err)
	}
	var awardsList = ConvertAwards(awards)
	return respx.SingleResp("request success", awardsList), nil
}

func ConvertAwards(awards []*model.Awards) []*types.PublicAwards {
	var awardsList []*types.PublicAwards
	for _, award := range awards {
		awardsList = append(awardsList, ConvertAward(award))
	}
	return awardsList
}

func ConvertAward(award *model.Awards) *types.PublicAwards {
	return &types.PublicAwards{
		Id:    award.Id,
		Uuid:  award.Uuid,
		Grade: award.Grade,
		Name:  award.Name,
		Image: award.Image,
		Price: award.Price,
		IsWin: award.IsWin > 0,
	}
}
