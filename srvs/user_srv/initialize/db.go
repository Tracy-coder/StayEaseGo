package initialize

import (
	"StayEaseGo/srvs/user_srv/global"
	"StayEaseGo/srvs/user_srv/model"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitSql() {
	cfg := global.GlobalServerConfig.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("can not connect to mysql: %s", dsn)
	}
	global.GlobalSqlClient = db

	err = global.GlobalSqlClient.AutoMigrate(&model.User{}, &model.UserAuth{})
	if err != nil {
		log.Fatal("auto migrate failed")
	}
}

func InitRedis() {
	cfg := global.GlobalServerConfig.RedisInfo
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Connect to redis client failed, err: %v\n", err)
	}
	global.GlobalRedisClient = rdb

}
