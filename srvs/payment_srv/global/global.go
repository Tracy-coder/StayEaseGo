package global

import (
	"StayEaseGo/srvs/payment_srv/config"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

const PaymentServiceConfigPath = "./srvs/payment_srv/config/config.yaml"

var (
	GlobalServerConfig config.ServerConfig
	GlobalConsulConfig config.ConsulConfig
)

var (
	GlobalSqlClient *gorm.DB
)

var (
	RocketMQProducer rocketmq.Producer
)

var (
	KafkaProducer *kafka.Writer
)
