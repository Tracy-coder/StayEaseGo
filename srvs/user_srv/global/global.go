package global

import (
	"StayEaseGo/srvs/user_srv/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const UserServiceConfigPath = "./srvs/user_srv/config/config.yaml"

var (
	GlobalServerConfig config.ServerConfig
	GlobalConsulConfig config.ConsulConfig
)

var (
	GlobalRedisClient *redis.Client
	GlobalSqlClient   *gorm.DB
)
