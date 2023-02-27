package users

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUsersLogic {
	return &DelUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelUsersLogic) DelUsers(req *types.DelUsersReq) (any, error) {
	if err := l.svcCtx.UsersModel.Delete(l.ctx, req.Id); err != nil {
		l.Logger.Errorf("删除用户失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
