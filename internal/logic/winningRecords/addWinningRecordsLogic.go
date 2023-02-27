/*
 * @Author: licat
 * @Date: 2023-02-20 23:49:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-21 13:09:27
 * @Description: licat233@gmail.com
 */
package winningRecords

import (
	"context"
	"net/http"

	"luckydraw-backend/common/utils"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddWinningRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAddWinningRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *AddWinningRecordsLogic {
	return &AddWinningRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *AddWinningRecordsLogic) AddWinningRecords(req *types.AddWinningRecordsReq) (resp *types.BaseResp, err error) {
	err = l.InsertWinningRecords(req)
	return
}

func (l *AddWinningRecordsLogic) InsertWinningRecords(req *types.AddWinningRecordsReq) (err error) {

	ip := utils.RemoteIp(l.r)

	platform, err := utils.RemotePlatform(l.r)
	if err != nil {
		platform = err.Error()
	}

	_, err = l.svcCtx.WinningRecordsModel.Insert(l.ctx, &model.WinningRecords{
		UserId:     req.UserId,
		AwardId:    req.AwardId,
		ActivityId: req.ActivityId,
		Ip:         ip,
		Platform:   platform,
	})
	return
}
