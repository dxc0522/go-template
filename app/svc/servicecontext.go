package svc

import (
	"gorm.io/gorm"

	"github.com/go-template/app/config"
	"github.com/go-template/pkg/provide"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) (*ServiceContext, error) {
	db, err := provide.Postgres(c.Name, c.Postgres.DSN)
	if err != nil {
		return nil, err
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}, nil
}
