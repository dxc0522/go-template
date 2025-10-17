package provide

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type PingResponse = struct {
	Time    string `json:"time"`
	Version string `json:"version"`
}

func ServerSetup(server *rest.Server, c rest.RestConf) {
	currentTime := time.Now().Format(time.DateTime)
	version := getVersion()
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/api/%s/ping", c.Name),
		Handler: func(w http.ResponseWriter, r *http.Request) {
			httpx.OkJson(w, PingResponse{
				Time:    currentTime,
				Version: version,
			})
		},
	})
	for _, route := range server.Routes() {
		log.Printf("%s - %s", route.Method, route.Path)
	}
	log.Printf("Starting server at http://%s:%d\n\n\n", c.Host, c.Port)
	server.Start()
}
func getVersion() string {
	// 首先检查环境变量 VERSION
	if version := os.Getenv("VERSION"); version != "" {
		return version
	}

	// 如果没有设置 VERSION 环境变量，则尝试获取 Git 分支名
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		// 如果无法获取分支名，返回默认值
		return "unknown"
	}

	return strings.TrimSpace(string(output))
}
