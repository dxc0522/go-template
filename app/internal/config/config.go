package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth     Auth
	DBConfig DBConfig
}

type Auth struct {
	AccessSecret string
	AccessExpire int64
}
type DBConfig struct {
	DriverName string
	Database   string
}
