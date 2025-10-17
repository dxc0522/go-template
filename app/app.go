package main

import (
	"flag"
	"os"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/go-template/app/config"
	"github.com/go-template/app/handler"
	"github.com/go-template/app/svc"
	"github.com/go-template/pkg/provide"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	_ = logx.SetUp(logx.LogConf{
		ServiceName: c.Name,
		Stat:        false,
		Level:       "info",
	})
	if c.ServerPort != 0 {
		c.RestConf.Port = c.ServerPort
	}
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx, err := svc.NewServiceContext(c)
	if err != nil {
		logx.Error(err)
		os.Exit(0)
	}

	handler.RegisterHandlers(server, ctx)

	provide.ServerSetup(server, c.RestConf)
}
