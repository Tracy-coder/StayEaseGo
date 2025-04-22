package global

import (
	"StayEaseGo/apis/order_web/biz/config"
	pb "StayEaseGo/srvs/order_srv/proto/gen"
)

const OrderServiceConfigPath = "./apis/order_web/biz/config/config.yaml"

var (
	GlobalServerConfig config.ServerConfig
	GlobalConsulConfig config.ConsulConfig
)

var (
	OrderSrvClient pb.OrderClient
)
