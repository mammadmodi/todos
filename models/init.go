package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"mods/todos/conf"
)

func Init(conf *conf.DBConf) *gorm.DB {
	db, err := gorm.Open("mysql", concatDBArgs(conf))
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Task{})
	log.Println("migration done.")

	return db
}

func concatDBArgs(dbConf *conf.DBConf) string {
	return dbConf.MysqlUser + ":" + dbConf.MysqlPass + "@/" + dbConf.DBName + "?charset=" + dbConf.MysqlCharSet
}
