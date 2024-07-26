package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth     Auth
	DBConfig DBConfig
}

type Auth struct {
	AccessSecret string `json:",env=AuthAccessSecret"`
	AccessExpire int64  `json:",env=AuthAccessExpire"`
}
type DBConfig struct {
	DriverName string `json:",env=DBDriverName"`
	Database   string `json:",env=DBDatabase"`
}
