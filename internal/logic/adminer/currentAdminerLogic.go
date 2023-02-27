/*
 * @Author: licat
 * @Date: 2023-02-22 17:54:34
 * @LastEditors: licat
 * @LastEditTime: 2023-02-22 19:46:34
 * @Description: licat233@gmail.com
 */
package adminer

import (
	"context"
	"errors"

	"luckydraw-backend/common/ctxdata"
	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/logic/tools"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentAdminerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentAdminerLogic {
	return &CurrentAdminerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentAdminerLogic) CurrentAdminer() (any, error) {
	adminerId := ctxdata.GetAdminerIdFromCtx(l.ctx)
	if adminerId == 0 {
		return nil, errorx.AuthError(errors.New("未登录"))
	}
	adminer, err := l.svcCtx.AdminerModel.FindOne(l.ctx, adminerId)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.AuthError(errors.New("用户不存在"))
	}
	data := tools.AdminerToResp(adminer)
	return respx.New(data), nil
}
