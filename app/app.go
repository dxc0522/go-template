package main

import (
	"flag"
	"fmt"
	"github.com/go-template/app/internal/config"
	"github.com/go-template/app/internal/handler"
	"github.com/go-template/app/internal/svc"
	"github.com/go-template/common/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"log"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/app-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(JwtUnauthorizedResult))
	defer server.Stop()

	ctx, err := svc.NewServiceContext(c)
	if err != nil {
		logx.Error(err)
		os.Exit(0)
	}
	handler.RegisterHandlers(server, ctx)

	for _, route := range server.Routes() {
		log.Printf("%s - %s", route.Method, route.Path)
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// JwtUnauthorizedResult jwt验证失败的回调
func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	logx.Info("jwt unauthorized", err)
	httpx.WriteJson(w, http.StatusUnauthorized, response.Body{Code: http.StatusUnauthorized, Msg: http.StatusText(http.StatusUnauthorized)})
}
