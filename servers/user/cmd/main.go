package main

import (
	"github.com/Zhan9Yunhua/blog-svr/servers/user/config"
	"github.com/Zhan9Yunhua/blog-svr/servers/user/server"
	"net/http"

	_ "github.com/Zhan9Yunhua/blog-svr/servers/user/config"
	"github.com/Zhan9Yunhua/blog-svr/servers/user/etcd"
	"github.com/Zhan9Yunhua/blog-svr/servers/user/logger"
	"github.com/Zhan9Yunhua/blog-svr/servers/user/middleware"
	"github.com/Zhan9Yunhua/blog-svr/servers/user/service"
)

func main() {
	lg := logger.NewLogger()

	etcdClient := etcd.NewEtcd()

	register := etcd.Register(etcdClient, lg)
	defer register.Deregister()

	var ucenterSvc service.UcenterServiceInterface
	ucenterSvc = service.UcenterService{}
	ucenterSvc = middleware.InstrumentingMiddleware()(ucenterSvc)

	mux := http.NewServeMux()
	mux.Handle("/svc/user/v1/", service.MakeHandler(ucenterSvc, lg))

	server.RunServer(mux,config.GetConfig().ServerPort)
}
