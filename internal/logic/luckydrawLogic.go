/*
 * @Author: licat
 * @Date: 2023-02-21 00:17:14
 * @LastEditors: licat
 * @LastEditTime: 2023-02-21 13:06:06
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"luckydraw-backend/common/ctxdata"
	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	awardslogic "luckydraw-backend/internal/logic/awards"
	"luckydraw-backend/internal/logic/winningRecords"

	"github.com/zeromicro/go-zero/core/logx"
)

type LuckydrawLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewLuckydrawLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *LuckydrawLogic {
	return &LuckydrawLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

// 要求指定用户才能抽，且只能抽中一次
func (l *LuckydrawLogic) Luckydraw(req *types.LuckydrawReq) (any, error) {
	var errorZeroCount = errorx.New("Sorry！你的抽獎次數為0，可聯絡客服獲取")
	// 一、先判断该用户是否已经注册
	userId := ctxdata.GetUidFromCtx(l.ctx)
	activityId := ctxdata.GetAidFromCtx(l.ctx)

	passport := strings.TrimSpace(req.Passport)
	activityUuid := strings.TrimSpace(req.ActivityUuid)
	if passport == "" && userId < 1 {
		return nil, errorZeroCount
	}
	if activityId < 1 && activityUuid == "" {
		return nil, errorx.New("錯誤的請求參數")
	}
	var err error
	var activity *model.Activity
	//如果activityId不存在，则通過activityUuid查询
	if activityId < 1 {
		activity, err = l.svcCtx.ActivityModel.FindOneByUuid(l.ctx, activityUuid)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询活动失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
	} else {
		//否则通过activityId查询
		activity, err = l.svcCtx.ActivityModel.FindOne(l.ctx, activityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询活动失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
	}
	if err == model.ErrNotFound || activity == nil {
		return nil, errorx.New("活動已不存在")
	}

	if activity.Status != 1 {
		return nil, errorx.New("Sorry！此活動已結束，感謝你的關注，請添加官方客服，獲取最新活動。")
	}

	activityId = activity.Id

	var user *model.Users
	//如果userId不存在，则通过passport查询
	if userId < 1 {
		user, err = l.svcCtx.UsersModel.FindsByPassportAndActivityId(l.ctx, passport, activityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询用户失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
	} else {
		user, err = l.svcCtx.UsersModel.FindsByIdAndActivityId(l.ctx, userId, activityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询用户失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
	}

	if err == model.ErrNotFound || user == nil {
		//如果没有注册，那就提示抽奖次数为0
		return nil, errorZeroCount
	}

	userId = user.Id

	if user.Count >= user.Total {
		return nil, errorx.New("Sorry！你的抽獎次數已用完，請聯絡客服獲取", "")
	}

	//获取其可抽中的奖品
	var availableAwardsIds []int64
	if user != nil {
		availableAwardsIds = getAvailableAwardsId(user.AvailableAwards)
	}

	//如果没有可抽中的奖品
	if len(availableAwardsIds) == 0 {
		//找个理由搪塞过去
		return nil, errorx.New("🔥🔥當前活動太火爆，伺服器擁堵，請稍後再重試...")
	}

	//获取奖品列表
	awards, err := l.svcCtx.AwardsModel.FindsByActivityId(l.ctx, activityId)
	if err != nil {
		l.Logger.Errorf("查询奖品列表失败，err:%v", err)
		return nil, errorx.InternalError(err)
	}
	if len(awards) == 0 {
		//找个理由搪塞过去
		return nil, errorx.New("🔥🔥當前活動太火爆，伺服器擁堵，請稍後再重試...")
	}

	// isRegistered := userId > 0

	//非註冊用戶，不让中奖
	// mustFail := !isRegistered
	award, err := randomAward(awards, false, availableAwardsIds)
	if err != nil {
		l.Logger.Errorf("随机奖品失败，err:%v", err)
		return nil, errorx.InternalError(err)
	}

	//存入记录
	recordlogic := winningRecords.NewAddWinningRecordsLogic(l.ctx, l.svcCtx, l.r)
	err = recordlogic.InsertWinningRecords(&types.AddWinningRecordsReq{
		UserId:     userId,
		AwardId:    award.Id,
		ActivityId: activityId,
	})
	if err != nil {
		l.Logger.Errorf("存入记录失败，err:%v", err)
		return nil, errorx.InternalError(err)
	}

	//同时将奖品的抽奖次数+1
	if err = l.svcCtx.AwardsModel.CountAdd(l.ctx, 1, award.Id); err != nil {
		l.Logger.Errorf("更新奖品抽奖次数失败，err:%v", err)
		return nil, errorx.InternalError(err)
	}
	//用户的抽奖次数也加+1
	if err = l.svcCtx.UsersModel.CountAdd(l.ctx, 1, user.Id); err != nil {
		l.Logger.Errorf("更新用户抽奖次数失败，err:%v", err)
		return nil, errorx.InternalError(err)
	}

	publicAward := awardslogic.ConvertAward(award)
	return respx.SingleResp("request success", publicAward), nil
}

/**
 * randomAward  抽奖
 * mustFail bool 是否必须抽不中
 * awardsRange []int64 指定其可抽中的奖品有哪些，如果为空，则默认抽中所有奖品，如果不为空，则按照指定的奖品有哪些抽中
 */
func randomAward(awards []*model.Awards, mustFail bool, awardsRange []int64) (*model.Awards, error) {
	//注意：考虑某个傻子设置幸运值为负数，负数或者0的prize不参与抽奖，因为逻辑上完全不可能抽中
	if len(awards) == 0 {
		return nil, errors.New("the awards list is empty")
	}
	var totalProb int64
	var prizes []*model.Awards
	for _, prize := range awards {
		//必须不中奖，不赢，概率大于等于0
		if mustFail && (prize.IsWin <= 0) && (prize.Prob >= 0) {
			return prize, nil
		}
		//如果运气值 <=0 ，则不参与抽奖
		if prize.Prob <= 0 {
			continue
		}
		if len(awardsRange) > 0 {
			//指定奖品
			for _, awardRange := range awardsRange {
				if prize.Id == awardRange {
					totalProb += prize.Prob
					prizes = append(prizes, prize)
					break
				}
			}
		} else {
			//默认全部
			totalProb += prize.Prob
			prizes = append(prizes, prize)
		}
	}

	//如果没有可抽的奖品
	if len(prizes) == 0 {
		return nil, errors.New("no available awards")
	}

	//如果全都不可能抽中
	if totalProb == 0 {
		return nil, errors.New("randomPrize Error, totalLuckValue = 0")
	}

	//如果只有一个，则不用再随机了，直接返回
	if len(prizes) == 1 {
		return prizes[0], nil
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(s1)
	random := rng.Int63n(totalProb)
	temp := random

	sortPrizes(prizes)

	//特殊情况，如果mustFail为true，则返回概率最大的那个奖品
	if mustFail {
		return prizes[len(prizes)-1], nil
	}

	//抽取算法
	for _, prize := range prizes {
		if random < prize.Prob {
			return prize, nil
		}
		random -= prize.Prob
	}
	return nil, fmt.Errorf("randomPrize Error, totalProb=%d, random=%d", totalProb, temp)
}

// 升序排序
func sortPrizes(awards []*model.Awards) {
	sort.Slice(awards, func(i, j int) bool {
		return awards[i].Prob < awards[j].Prob
	})
}

func getAvailableAwardsId(str string) []int64 {
	var ids []int64
	for _, s := range strings.Split(str, ",") {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			continue
		}
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}
