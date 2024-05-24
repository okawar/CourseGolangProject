package sql

import (
	"fmt"
	"golang_pr/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB
var migrations = make([]func(), 0)

func init() {
	var err error
	database, err = gorm.Open(sqlite.Open(config.Get().Server.Workdir + "/sql/NewStorage.sqlite3"))
	if err != nil {
		fmt.Println(config.Get().Server.Workdir)
		//panic(err)
	}
}

func AddMigrations(mF func()) {
	migrations = append(migrations, mF)
}

func Migrate() {
	for _, f := range migrations {
		f()
	}
}

func GetDB() *gorm.DB {
	return database
}

func GetMigrations() []func() {
	return migrations
}
