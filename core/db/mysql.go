package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/yii2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Mysql connect success！")
	Mysql = db
}
