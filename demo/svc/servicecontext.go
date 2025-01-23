package svc

import (
	"github.com/go-template/demo/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) (*ServiceContext, error) {
	return &ServiceContext{
		Config: c,
	}, nil
}
