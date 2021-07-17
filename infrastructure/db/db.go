package db

import (
	"github.com/cogny/go_verdao/application/model"
	"github.com/jinzhu/gorm"

	_ "gorm.io/driver/sqlite"
)

func ConnectDB() *gorm.DB {
	dsn := ":memory:"
	db, err := gorm.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	db.AutoMigrate(&model.Result{})
	return db
}
