/*
 * @Author: licat
 * @Date: 2023-02-20 23:49:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-22 20:15:59
 * @Description: licat233@gmail.com
 */
package main

import (
	"flag"
	"fmt"
	"net/http"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/internal/config"
	"luckydraw-backend/internal/handler"
	"luckydraw-backend/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/luckydraw-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf, rest.WithCors())

	defer server.Stop()

	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/api/static/:file",
		Handler: http.StripPrefix("/api/static/", http.FileServer(http.Dir("./static"))).ServeHTTP,
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//错误处理
	httpx.SetErrorHandlerCtx(errorx.ErrorResponseHandlerCtx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
