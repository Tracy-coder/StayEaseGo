package initialize

import (
	"StayEaseGo/srvs/homestay_srv/global"
	"StayEaseGo/srvs/homestay_srv/model"
	"fmt"

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

	err = global.GlobalSqlClient.AutoMigrate(&model.Homestay{})
	if err != nil {
		log.Fatal("auto migrate failed")
	}
}
