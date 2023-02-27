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

type PutAdminerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutAdminerLogic {
	return &PutAdminerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutAdminerLogic) PutAdminer(req *types.PutAdminerReq) (any, error) {
	if err := validateAccessValue(req.Access); err != nil {
		return nil, errorx.New(err.Error())
	}
	in := &model.Adminer{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Access:   req.Access,
		IsSuper:  req.IsSuper,
	}

	err := l.svcCtx.AdminerModel.Update(l.ctx, in)
	if err != nil {
		l.Logger.Error("failed to update adminer, error: ", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
