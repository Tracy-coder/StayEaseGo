package main

import (
	"StayEaseGo/srvs/mq/global"
	"StayEaseGo/srvs/mq/handler"
	"StayEaseGo/srvs/mq/initialize"
	"StayEaseGo/srvs/mq/model"
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/segmentio/kafka-go"
)

var OrderSuccessNotifyUserConsumer *kafka.Reader
var ThirdPaymentUpdatePayStatusNotifyConsumer *kafka.Reader

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitOrderSrv()
	initialize.InitUserSrv()
	log.Info(global.GlobalServerConfig)
	OrderSuccessNotifyUserConsumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  global.GlobalServerConfig.KafkaCfg.Brokers,
		GroupID:  "0",
		Topic:    model.JobTypeOrderSuccessNotifyUser,
		MaxBytes: 10e6, // 10MB
	})
	ThirdPaymentUpdatePayStatusNotifyConsumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  global.GlobalServerConfig.KafkaCfg.Brokers,
		GroupID:  "0",
		Topic:    model.JobTypeThirdPaymentUpdatePayStatusNotify,
		MaxBytes: 10e6, // 10MB
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go ConsumeThirdPaymentUpdatePayStatusNotifyMessages(ctx)
	go ConsumeOrderSuccessNotifyUserMessages(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signalChan:
		log.Println("Signal received, shutting down...")
		cancel()
	}
}

func ConsumeThirdPaymentUpdatePayStatusNotifyMessages(ctx context.Context) error {
	for {
		msg, err := ThirdPaymentUpdatePayStatusNotifyConsumer.ReadMessage(ctx)
		if err != nil {
			log.Error(err)
		}
		var payload model.ThirdPaymentUpdatePayStatusNotifyMessage
		json.Unmarshal(msg.Value, &payload)
		log.Info("payment update: ", payload)
		err = handler.ThirdPaymentUpdatePayStatusNotifyHandler(payload)
		if err != nil {
			log.Error(err)
		}
	}
}

func ConsumeOrderSuccessNotifyUserMessages(ctx context.Context) error {
	for {
		msg, err := OrderSuccessNotifyUserConsumer.ReadMessage(ctx)
		if err != nil {
			log.Error(err)
		}
		var payload model.OrderSuccessNotifyUserMessage
		json.Unmarshal(msg.Value, &payload)
		log.Info("order success: ", payload)
		err = handler.OrderSuccessNotifyUserHandler(payload)
		if err != nil {
			log.Error(err)
		}
	}
}
