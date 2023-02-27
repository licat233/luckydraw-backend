/*
 * @Author: licat
 * @Date: 2023-02-22 17:15:15
 * @LastEditors: licat
 * @LastEditTime: 2023-02-22 20:18:18
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"errors"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test() (resp *types.BaseResp, err error) {
	resp = new(types.BaseResp)
	resp.Status = true
	resp.Message = "hello"
	err = errorx.ExternalError(errors.New("test error"))
	return
}
