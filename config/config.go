package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

var Mysql mysqlC
var Redice redisC

type mysqlC struct {
	Host     string
	Port     string
	User     string
	Password string
}

type redisC struct {
	Host     string
	Port     string
	User     string
	Password string
}

type config struct {
	Db struct {
		Mysql mysqlC
		Redis redisC
	}
}

func Load() {
	//读取配置信息
	c := config{}
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal("配置文件读取错误:" + err.Error())
	}

	Mysql = c.Db.Mysql

	log.Println(Mysql)
}
