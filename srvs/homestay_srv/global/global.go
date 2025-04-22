package global

import (
	"StayEaseGo/srvs/homestay_srv/config"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

const HomestayServiceConfigPath = "./srvs/homestay_srv/config/config.yaml"

var (
	GlobalServerConfig config.ServerConfig
	GlobalConsulConfig config.ConsulConfig
)

var (
	GlobalSqlClient *gorm.DB
	GlobalESClient  *elasticsearch.Client
)
