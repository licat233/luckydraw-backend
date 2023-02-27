package logic

import (
	"context"
	"net/http"

	"luckydraw-backend/common/respx"
	"luckydraw-backend/common/utils"
	"luckydraw-backend/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRemoteIpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetRemoteIpLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetRemoteIpLogic {
	return &GetRemoteIpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetRemoteIpLogic) GetRemoteIp() (any, error) {
	ip := utils.RemoteIp(l.r)
	return respx.New(ip), nil
}
