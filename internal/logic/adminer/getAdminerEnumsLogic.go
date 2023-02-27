package adminer

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminerEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerEnumsLogic {
	return &GetAdminerEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminerEnumsLogic) GetAdminerEnums(req *types.GetAdminerEnumsReq) (any, error) {
	list, _, err := l.svcCtx.AdminerModel.FindList(l.ctx, 0, 0, "", nil)
	if err != nil {
		l.Logger.Errorf("获取奖品列表失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	var listEnums []*types.Enum
	for _, v := range list {
		listEnums = append(listEnums, &types.Enum{
			Label: v.Username,
			Value: v.Id,
		})
	}
	return respx.New(listEnums), nil
}
