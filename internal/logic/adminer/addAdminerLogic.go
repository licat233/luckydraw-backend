package adminer

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAdminerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAdminerLogic {
	return &AddAdminerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAdminerLogic) AddAdminer(req *types.AddAdminerReq) (any, error) {
	if err := validateAccessValue(req.Access); err != nil {
		return nil, errorx.New(err.Error())
	}
	in := &model.Adminer{
		Id:       0,
		Username: req.Username,
		Password: req.Password,
		Access:   req.Access,
		IsSuper:  req.IsSuper,
	}
	_, err := l.svcCtx.AdminerModel.Insert(l.ctx, in)
	if err != nil {
		l.Logger.Error("failed to insert adminer, error: ", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
