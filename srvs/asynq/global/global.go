package global

import (
	"StayEaseGo/srvs/asynq/config"
	order_srv "StayEaseGo/srvs/order_srv/proto/gen"

	"github.com/hibiken/asynq"
)

const AsynqServiceConfigPath = "./srvs/asynq/config/config_docker.yaml"

var (
	AsynqServer *asynq.Server
)

var (
	GlobalServerConfig config.ServerConfig
	GlobalConsulConfig config.ConsulConfig
)
var (
	OrderSrvClient order_srv.OrderClient
)
