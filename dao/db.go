package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-learning/log"
)

var db *gorm.DB

func init() {
}

func dbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	return db
}

func GetDb() *gorm.DB {
	return dbConn("root", "root", "127.0.0.1", "spring_jpa", 3306)
}
