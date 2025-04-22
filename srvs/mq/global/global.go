package global

import (
	"StayEaseGo/srvs/mq/config"
	order_srv "StayEaseGo/srvs/order_srv/proto/gen"
	user_srv "StayEaseGo/srvs/user_srv/proto/gen"
)

const MQServiceConfigPath = "./srvs/mq/config/config.yaml"

var (
	GlobalServerConfig config.ServerConfig
	GlobalConsulConfig config.ConsulConfig
)

var (
	UserSrvClient  user_srv.UserClient
	OrderSrvClient order_srv.OrderClient
)
