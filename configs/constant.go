package configs

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	Cache       *redis.Client
	SecretKey   string
	Issuer      string
	ExpiresTime int
)
