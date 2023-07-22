/*
 * @Author: licat
 * @Date: 2023-02-20 23:49:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-22 14:37:32
 * @Description: licat233@gmail.com
 */
package svc

import (
	"log"
	"luckydraw-backend/common/mysqlx"
	"luckydraw-backend/internal/config"
	"luckydraw-backend/internal/middleware"
	"luckydraw-backend/model"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config             config.Config
	AuthMiddleware     rest.Middleware
	UserAuthMiddleware rest.Middleware

	CaptchaStore  base64Captcha.Store
	CaptchaExpire time.Duration

	UsersModel          model.UsersModel
	ActivityModel       model.ActivityModel
	AwardsModel         model.AwardsModel
	WinningRecordsModel model.WinningRecordsModel
	AdminerModel        model.AdminerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn, err := mysqlx.NewMysql(c.Mysql.DataSource)
	if err != nil {
		log.Fatalln("数据库初始化失败", c.Mysql.DataSource)
	}

	captchaExpire := time.Minute * 5

	return &ServiceContext{
		Config:              c,
		AuthMiddleware:      middleware.NewAuthMiddleware(c).Handle,
		UserAuthMiddleware:  middleware.NewUserAuthMiddleware().Handle,
		CaptchaStore:        base64Captcha.NewMemoryStore(10240, captchaExpire),
		CaptchaExpire:       captchaExpire,
		UsersModel:          model.NewUsersModel(sqlConn),
		ActivityModel:       model.NewActivityModel(sqlConn),
		AwardsModel:         model.NewAwardsModel(sqlConn),
		WinningRecordsModel: model.NewWinningRecordsModel(sqlConn),
		AdminerModel:        model.NewAdminerModel(sqlConn),
	}
}
