package db

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Redis *redis.Client
	Mysql *gorm.DB
)
