package main

import (
	"StayEaseGo/srvs/asynq/global"
	"StayEaseGo/srvs/asynq/handler"
	"StayEaseGo/srvs/asynq/initialize"
	"os"

	"StayEaseGo/srvs/asynq/model"

	"github.com/hibiken/asynq"
	log "github.com/sirupsen/logrus"
)

func main() {
	initialize.InitConfig()
	log.Debug(global.GlobalServerConfig)
	initialize.InitLogger()
	initialize.InitAsynqServer()
	initialize.InitOrderSrv()
	mux := asynq.NewServeMux()

	//scheduler job
	mux.Handle(model.JobTypeOrderAutoCancelNotify, handler.NewOrderAutoCancelHandler())
	if err := global.AsynqServer.Run(mux); err != nil {
		log.Errorf("run asynq server error: %s", err.Error())
		os.Exit(1)
	}

}
