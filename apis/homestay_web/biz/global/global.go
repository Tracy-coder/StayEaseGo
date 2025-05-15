package global

import (
	"StayEaseGo/apis/homestay_web/biz/config"
	pb "StayEaseGo/srvs/homestay_srv/proto/gen"
	user_srv "StayEaseGo/srvs/user_srv/proto/gen"
)

const HomestayServiceConfigPath = "./apis/homestay_web/biz/config/config_docker.yaml"

var (
	GlobalServerConfig config.ServerConfig
	GlobalConsulConfig config.ConsulConfig
)

var (
	HomestaySrvClient pb.HomestayClient
	UserSrvClient     user_srv.UserClient
)
