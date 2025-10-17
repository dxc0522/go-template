package provide

import (
	"errors"
	"fmt"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const migrationsDir = "./migrations"

func Mysql(name string, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connect Mysql Error, error=" + err.Error())
	} else {
		fmt.Println("Connect Mysql Success")
	}

	tableName := fmt.Sprintf("%s_goose_version", name)
	_ = goose.SetDialect("mysql")
	goose.SetTableName(tableName)
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := goose.Up(sqlDB, migrationsDir); errors.Is(err, goose.ErrNoMigrations) && !errors.Is(err, goose.ErrNoMigrationFiles) {
		return nil, fmt.Errorf("goose up failed: %w \n", err)
	}
	return db.Debug(), nil
}
