package users

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersEnumsLogic {
	return &GetUsersEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersEnumsLogic) GetUsersEnums(req *types.GetUsersEnumsReq) (any, error) {
	list, _, err := l.svcCtx.UsersModel.FindList(l.ctx, 0, 0, "", nil)
	if err != nil {
		l.Logger.Errorf("获取用户列表失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	var listEnums []*types.Enum
	for _, v := range list {
		listEnums = append(listEnums, &types.Enum{
			Label: v.Name,
			Value: v.Id,
		})
	}
	return respx.New(listEnums), nil
}
