package main

import (
	"github.com/Zhan9Yunhua/blog-svr/gateway/config"
	"github.com/Zhan9Yunhua/blog-svr/gateway/etcd"
	"github.com/Zhan9Yunhua/blog-svr/gateway/logger"
	"github.com/Zhan9Yunhua/blog-svr/gateway/middleware"
	"github.com/Zhan9Yunhua/blog-svr/gateway/router"
	"github.com/Zhan9Yunhua/blog-svr/gateway/server"
)

func main() {
	lg := logger.NewLogger()

	etcdClient := etcd.NewEtcd()

	r := router.NewRouter(lg)
	{
		r.Service("/svc/user", etcdClient)
		r.Post("/svc/user/login", middleware.CookieMiddleware())
		r.JetGet("/svc/user/{param}",  middleware.CookieMiddleware())
	}

	server.RunServer(config.GetConfig().ServerPort, r)
}
