package models

import (
	"github.com/jinzhu/gorm"
)

type DBConf struct {
	MysqlHost    string `default:"127.0.0.1" split_words:"true"`
	MysqlPort    string `default:"3306" split_words:"true"`
	MysqlUser    string `default:"root" split_words:"true"`
	MysqlPass    string `default:"123456" split_words:"true"`
	MysqlCharSet string `default:"utf8" split_words:"true"`
	DBName       string `default:"todos" split_words:"true"`
}

func InitDB(conf *DBConf) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", concatDBArgs(conf))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func concatDBArgs(dbConf *DBConf) string {
	return dbConf.MysqlUser + ":" + dbConf.MysqlPass + "@/" + dbConf.DBName + "?parseTime=true&charset=" + dbConf.MysqlCharSet
}
