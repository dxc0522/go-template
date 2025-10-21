package provide

import (
	"errors"
	"fmt"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres(name string, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connect Postgres Error, error=" + err.Error())
	} else {
		fmt.Println("Connect Postgres Success")
	}

	tableName := fmt.Sprintf("%s_goose_version", name)
	_ = goose.SetDialect("postgres")
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
