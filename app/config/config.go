package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	ServerPort int    `json:",env=PORT,default=8080"`
	AppMode    string `json:",env=APP_MODE,default=LOCAL"`
	Postgres   struct {
		DSN string `json:",env=POSTGRES_DSN"`
	}
}
