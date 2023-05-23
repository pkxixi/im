package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"im/config"
)

var (
	ServiceConfig *config.ServiceConfig
	DB            *gorm.DB
	Redis         *redis.Client
)
