package initialize

import (
	"StayEaseGo/srvs/order_srv/global"
	"fmt"

	"github.com/hibiken/asynq"
)

func InitAsynqClient() {
	global.AsynqClient = asynq.NewClient(asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%d", global.GlobalServerConfig.RedisInfo.Host, global.GlobalServerConfig.RedisInfo.Port), Password: global.GlobalServerConfig.RedisInfo.Password})
}
