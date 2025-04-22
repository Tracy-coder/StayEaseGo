package initialize

import (
	"StayEaseGo/srvs/asynq/global"
	"fmt"

	"github.com/hibiken/asynq"
)

func InitAsynqServer() {
	global.AsynqServer = asynq.NewServer(
		asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%d", global.GlobalServerConfig.RedisInfo.Host, global.GlobalServerConfig.RedisInfo.Port), Password: global.GlobalServerConfig.RedisInfo.Password},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
			Concurrency: 20, //max concurrent process job task num
		},
	)
}
