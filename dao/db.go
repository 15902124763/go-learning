package dao

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"go-learning/log"
	_"github.com/go-sql-driver/mysql"
)

var db *gorm.DB


func init() {
}

func dbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser,Password, Host, Port, Db )
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	return db
}

func GetDb() *gorm.DB {
	return dbConn("root", "root", "127.0.0.1", "spring_jpa", 3306);
}