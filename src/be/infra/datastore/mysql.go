package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Conn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	return db
}

func GetDBConfig() (string, string) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := ""
	DBNAME := "anipark"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONFIG := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

	return DBMS, CONFIG
}
