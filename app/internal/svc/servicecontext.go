package svc

import (
	"fmt"
	"github.com/go-template/app/internal/config"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) (*ServiceContext, error) {
	db, err := gorm.Open(mysql.Open(c.DBConfig.Database), &gorm.Config{})
	if err != nil {
		panic("Connect Mysql Error, error=" + err.Error())
	} else {
		fmt.Println("Connect Mysql Success")
	}

	tableName := fmt.Sprintf("%s_goose_version", c.Name)
	//goose.SetBaseFS(os.DirFS("./app/migrations"))
	_ = goose.SetDialect("mysql")
	goose.SetTableName(tableName)
	sqlDB, _ := db.DB()
	err = goose.Up(sqlDB, "./migrations")
	if err != nil {
		return nil, err
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}, nil
}
