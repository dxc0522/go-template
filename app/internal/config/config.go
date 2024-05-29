package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth Auth
}

type Auth struct {
	AccessSecret string
	AccessExpire int64
}
type DBConfig struct {
	MySql struct {
		Database string
	}
}
