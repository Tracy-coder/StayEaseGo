package initialize

import (
	"StayEaseGo/srvs/order_srv/global"
	"time"

	mq "StayEaseGo/srvs/mq/model"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func InitMQ() {
	var err error
	global.KafkaProducer = &kafka.Writer{
		Addr:                   kafka.TCP(global.GlobalServerConfig.KafkaCfg.Brokers...),
		Topic:                  mq.JobTypeOrderSuccessNotifyUser,
		Balancer:               &kafka.LeastBytes{},
		WriteTimeout:           10 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}
	if err != nil {
		log.Fatal("producer init error:", err)
	}
}
