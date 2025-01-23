package provide

import (
	"fmt"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
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
	if _, err := os.Stat(migrationsDir); os.IsExist(err) {
		sqlDB, _ := db.DB()
		err = goose.Up(sqlDB, "./migrations")
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("GORM ERROR: Not Found Migrations Folder!")
	}
	return db.Debug(), nil
}
