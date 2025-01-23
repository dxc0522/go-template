package provide

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"log"
	"net/http"
	"os"
	"time"
)

type PingResponse = struct {
	Time    string `json:"time"`
	Version string `json:"version"`
}

func Context(server *rest.Server, c rest.RestConf) {
	currentTime := time.DateTime
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/api/%s/ping", c.Name),
		Handler: func(w http.ResponseWriter, r *http.Request) {
			httpx.OkJson(w, PingResponse{
				Time:    currentTime,
				Version: os.Getenv("VERSION"),
			})
		},
	})
	for _, route := range server.Routes() {
		log.Printf("%s - %s", route.Method, route.Path)
	}
	fmt.Printf("Starting server at http://%s:%d\n\n\n", c.Host, c.Port)
}
