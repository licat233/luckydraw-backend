/*
 * @Author: licat
 * @Date: 2023-02-21 12:42:08
 * @LastEditors: licat
 * @LastEditTime: 2023-02-22 17:55:16
 * @Description: licat233@gmail.com
 */
package ctxdata

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	CtxKeyJwtUserId     = "userId"
	CtxKeyJwtAdminerId  = "adminerId"
	CtxKeyJwtActivityId = "activityId"
	CtxKeyJwtIp         = "ip"
	CtxKeyJwtPlatform   = "palatform"
)

// 获取用户id
func GetUidFromCtx(ctx context.Context) int64 {
	var user_id int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			user_id = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err: %+v", err)
		}
	}
	return user_id
}

func GetAdminerIdFromCtx(ctx context.Context) int64 {
	var adminer_id int64
	if jsonAdminerId, ok := ctx.Value(CtxKeyJwtAdminerId).(json.Number); ok {
		if int64AdminerId, err := jsonAdminerId.Int64(); err == nil {
			adminer_id = int64AdminerId
		} else {
			logx.WithContext(ctx).Errorf("GetAdminerIdFromCtx err: %+v", err)
		}
	}
	return adminer_id
}

// 获取活动id
func GetAidFromCtx(ctx context.Context) int64 {
	var activity_id int64
	if jsonAid, ok := ctx.Value(CtxKeyJwtActivityId).(json.Number); ok {
		if int64Aid, err := jsonAid.Int64(); err == nil {
			activity_id = int64Aid
		} else {
			logx.WithContext(ctx).Errorf("GetAidFromCtx err: %+v", err)
		}
	}
	return activity_id
}

// 获取ip
func GetIpFromCtx(ctx context.Context) string {
	return fmt.Sprint(ctx.Value(CtxKeyJwtIp))
}

// 获取platform
func GetPlatformFromCtx(ctx context.Context) string {
	return fmt.Sprint(ctx.Value(CtxKeyJwtPlatform))
}
