package global

import (
	homestay "StayEaseGo/srvs/homestay_srv/proto/gen"
	"StayEaseGo/srvs/order_srv/config"

	"github.com/hibiken/asynq"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

const OrderServiceConfigPath = "./srvs/order_srv/config/config_docker.yaml"

var (
	GlobalServerConfig config.ServerConfig
	GlobalConsulConfig config.ConsulConfig
)

var (
	GlobalSqlClient *gorm.DB
)

var (
	HomestaySrvClient homestay.HomestayClient
)

var (
	KafkaProducer *kafka.Writer
)

var (
	AsynqClient *asynq.Client
)
