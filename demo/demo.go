package main

import (
	"flag"
	"github.com/go-template/common/pkg/provide"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"github.com/go-template/demo/config"
	"github.com/go-template/demo/handler"
	"github.com/go-template/demo/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/demo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	logx.SetUp(logx.LogConf{
		ServiceName: c.Name,
		Stat:        false,
	})

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx, err := svc.NewServiceContext(c)
	if err != nil {
		logx.Error(err)
		os.Exit(0)
	}

	handler.RegisterHandlers(server, ctx)

	provide.Context(server, c.RestConf)
	server.Start()
}
