package adminer

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

type GetAdminerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerListLogic {
	return &GetAdminerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminerListLogic) GetAdminerList(req *types.GetAdminerListReq) (any, error) {
	in := &model.Adminer{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Access:   req.Access,
		IsSuper:  req.IsSuper,
	}
	data, total, err := l.svcCtx.AdminerModel.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
	if err != nil {
		l.Logger.Errorf("获取管理员列表失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	var list []*types.Adminer
	for _, v := range data {
		list = append(list, tools.AdminerToResp(v))
	}
	return respx.NewListData(list, total, req.PageSize, req.Page), nil
}
